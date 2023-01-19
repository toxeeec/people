package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

const (
	SelectUser = "SELECT user_profile.user_id, handle, following, followers FROM user_profile"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.User {
	return &userRepo{db}
}

func (r *userRepo) GetID(handle string) (uint, error) {
	const query = "SELECT user_id FROM user_profile WHERE handle = $1"
	var id uint
	if err := r.db.Get(&id, query, handle); err != nil {
		return 0, fmt.Errorf("User.GetID: %w", err)
	}
	return id, nil
}

func (r *userRepo) Create(au people.AuthUser) (people.User, error) {
	const query = "INSERT INTO user_profile(handle, hash) VALUES ($1, $2) RETURNING user_id, handle, following, followers"
	var u people.User
	if err := r.db.Get(&u, query, au.Handle, au.Password); err != nil {
		return u, fmt.Errorf("User.Create: %w", err)
	}
	return u, nil
}

func (r *userRepo) Delete(id uint) error {
	const query = "DELETE FROM user_profile WHERE user_id = $1"
	if _, err := r.db.Exec(query, id); err != nil {
		return fmt.Errorf("User.Delete: %w", err)
	}
	return nil
}

func (r *userRepo) GetHash(id uint) (string, error) {
	const query = "SELECT hash FROM user_profile WHERE user_id = $1"
	var h string
	if err := r.db.Get(&h, query, id); err != nil {
		return "", fmt.Errorf("User.GetHash: %w", err)
	}
	return h, nil
}

func (r *userRepo) Get(id uint) (people.User, error) {
	const query = SelectUser + " WHERE user_id = $1"
	var u people.User
	if err := r.db.Get(&u, query, id); err != nil {
		return u, fmt.Errorf("User.Get: %w", err)
	}
	return u, nil
}

func (r *userRepo) List(ids []uint) ([]people.User, error) {
	if len(ids) == 0 {
		return []people.User{}, nil
	}
	const query = SelectUser + " WHERE user_id = $1"
	q, args, err := NewQuery(SelectUser).Where("user_id IN (?)", ids).Build()
	if err != nil {
		return nil, fmt.Errorf("User.List: %w", err)
	}
	us := make([]people.User, len(ids))
	if err := r.db.Select(&us, q, args...); err != nil {
		return nil, fmt.Errorf("User.List: %w", err)
	}
	return us, nil
}

func (r *userRepo) ListMatches(query string, p pagination.ID) ([]people.User, error) {
	q, args, err := NewQuery(SelectUser).
		Where("handle ILIKE ?", "%"+query+"%").
		Paginate(p, "user_id", "?").
		Build()
	if err != nil {
		return nil, fmt.Errorf("User.ListMatches: %w", err)
	}
	us := make([]people.User, p.Limit)
	if err := r.db.Select(&us, q, args...); err != nil {
		return nil, fmt.Errorf("User.ListMatches: %w", err)
	}
	return us, nil
}

func (r *userRepo) Update(userID uint, handle string) (people.User, error) {
	const query = "UPDATE user_profile SET handle = $1 WHERE user_id = $2 RETURNING user_id, handle, following, followers"
	var u people.User
	if err := r.db.Get(&u, query, handle, userID); err != nil {
		return people.User{}, fmt.Errorf("User.Update: %w", err)
	}
	return u, nil
}
