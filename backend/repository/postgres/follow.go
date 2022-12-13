package postgres

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type followRepo struct {
	db *sqlx.DB
}

func NewFollowRepository(db *sqlx.DB) repository.Follow {
	return &followRepo{db}
}

func (s *followRepo) GetStatusFollowing(srcID, userID uint) bool {
	const query = "SELECT EXISTS(SELECT 1 FROM follower WHERE follower_id = $1 AND user_id = $2)"
	var f bool
	s.db.Get(&f, query, srcID, userID)
	return f
}

func (s *followRepo) GetStatusFollowed(srcID, userID uint) bool {
	const query = "SELECT EXISTS(SELECT 1 FROM follower WHERE user_id = $1 AND follower_id = $2)"
	var f bool
	s.db.Get(&f, query, srcID, userID)
	return f
}

func (r *followRepo) Create(targetID, id uint) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("Follow.Create: %w", err)
	}
	defer tx.Rollback()
	const query = "INSERT INTO follower(user_id, follower_id) VALUES ($1, $2)"
	_, err = tx.Exec(query, targetID, id)
	if err != nil {
		var e *pq.Error
		if errors.As(err, &e) {
			if e.Constraint == "different_user" {
				return fmt.Errorf("Follow.Create: %w", repository.ErrSameUser)
			}
			if e.Constraint == "follower_pkey" {
				return fmt.Errorf("Follow.Create: %w", repository.ErrAlreadyFollowed)
			}
		}
		return fmt.Errorf("Follow.Create: %w", err)
	}

	const incrementFollowing = "UPDATE user_profile SET following = following + 1 WHERE user_id = $1"
	_, err = tx.Exec(incrementFollowing, id)
	if err != nil {
		return fmt.Errorf("Follow.Create: %w", err)
	}
	const incrementFollowers = "UPDATE user_profile SET followers = followers + 1 WHERE user_id = $1"
	_, err = tx.Exec(incrementFollowers, targetID)
	if err != nil {
		return fmt.Errorf("Follow.Create: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Follow.Create: %w", err)
	}
	return nil
}

func (r *followRepo) Delete(targetID, id uint) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("Follow.Delete: %w", err)
	}
	defer tx.Rollback()
	const query = "DELETE FROM follower WHERE user_id = $1 AND follower_id = $2 RETURNING user_id"
	err = tx.Get(new(uint), query, targetID, id)
	if err != nil {
		return fmt.Errorf("Follow.Delete: %w", err)
	}

	const decrementFollowing = "UPDATE user_profile SET following = following - 1 WHERE user_id = $1"
	_, err = tx.Exec(decrementFollowing, id)
	if err != nil {
		return fmt.Errorf("Follow.Delete: %w", err)
	}
	const decrementFollowers = "UPDATE user_profile SET followers = followers - 1 WHERE user_id = $1"
	_, err = tx.Exec(decrementFollowers, targetID)
	if err != nil {
		return fmt.Errorf("Follow.Delete: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Follow.Delete: %w", err)
	}
	return nil
}

func (r *followRepo) ListStatusFollowing(srcIDs []uint, userID uint) (map[uint]struct{}, error) {
	if len(srcIDs) == 0 {
		return make(map[uint]struct{}), nil
	}
	const query = "SELECT follower_id FROM follower"
	q, args, err := NewQuery(query).Where("follower_id IN (?)", srcIDs).Where("user_id = ?", userID).Build()
	if err != nil {
		return nil, fmt.Errorf("Follow.ListStatusFollowing: %w", err)
	}
	followingIDs := make([]uint, len(srcIDs))
	err = r.db.Select(&followingIDs, q, args...)
	if err != nil {
		return nil, fmt.Errorf("Follow.ListStatusFollowing: %w", err)
	}
	fss := make(map[uint]struct{}, len(srcIDs))
	for _, id := range followingIDs {
		fss[id] = struct{}{}
	}
	return fss, nil
}

func (r *followRepo) ListStatusFollowed(srcIDs []uint, userID uint) (map[uint]struct{}, error) {
	if len(srcIDs) == 0 {
		return make(map[uint]struct{}), nil
	}
	const query = "SELECT user_id FROM follower"
	q, args, err := NewQuery(query).Where("user_id IN (?)", srcIDs).Where("follower_id = ?", userID).Build()
	if err != nil {
		return nil, fmt.Errorf("Follow.ListStatusFollowed: %w", err)
	}
	followedIDs := make([]uint, len(srcIDs))
	err = r.db.Select(&followedIDs, q, args...)
	if err != nil {
		return nil, fmt.Errorf("Follow.ListStatusFollowed: %w", err)
	}
	fss := make(map[uint]struct{}, len(srcIDs))
	for _, id := range followedIDs {
		fss[id] = struct{}{}
	}
	return fss, nil
}

func (r *followRepo) ListFollowing(id uint, p pagination.ID) ([]people.User, error) {
	const paginationValue = "(SELECT followed_at FROM follower WHERE user_id = ? AND follower_id = ?)"
	q, args, err := NewQuery(SelectUser).
		Join("follower", "user_profile.user_id = follower.user_id").
		Where("follower_id = ?", id).
		Paginate(p, "followed_at", paginationValue, id).
		Build()
	if err != nil {
		return nil, fmt.Errorf("Follow.ListFollowing: %w", err)
	}
	var us []people.User
	if err := r.db.Select(&us, q, args...); err != nil {
		return nil, fmt.Errorf("Follow.ListFollowing: %w", err)
	}
	return us, nil
}

func (r *followRepo) ListFollowers(id uint, p pagination.ID) ([]people.User, error) {
	const paginationValue = "(SELECT followed_at FROM follower WHERE follower_id = ? AND user_id = ?)"
	q, args, err := NewQuery(SelectUser).
		Join("follower", "user_profile.user_id = follower.follower_id").
		Where("follower.user_id = ?", id).
		Paginate(p, "followed_at", paginationValue, id).
		Build()
	if err != nil {
		return nil, fmt.Errorf("Follow.ListFollowers: %w", err)
	}
	var us []people.User
	if err := r.db.Select(&us, q, args...); err != nil {
		return nil, fmt.Errorf("Follow.ListFollowers: %w", err)
	}
	return us, nil
}
