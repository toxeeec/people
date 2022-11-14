package user

import (
	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) people.UserService {
	return &service{db}
}

const (
	isFollowed = " EXISTS(SELECT 1 FROM follower WHERE follower_id = $1 AND user_id = user_profile.user_id) as is_followed"
	selectUser = "SELECT handle, following, followers," + isFollowing + "," + isFollowed + " FROM user_profile"
)

const (
	queryExists  = "SELECT EXISTS(SELECT 1 FROM user_profile WHERE handle = $1)"
	queryCreate  = "INSERT INTO user_profile(handle, hash) VALUES($1, $2) RETURNING user_id"
	queryDelete  = "DELETE FROM user_profile WHERE handle = $1"
	queryGetAuth = "SELECT user_id, handle, hash FROM user_profile WHERE handle = $1"
	queryGet     = selectUser + " WHERE handle = $2"
)

func (s *service) Exists(handle string) bool {
	var exists bool
	s.db.Get(&exists, queryExists, handle)
	return exists
}

// Create returns id of the created user.
func (s *service) Create(u people.AuthUser) (uint, error) {
	var id uint
	hash, err := u.Password.Hash()
	if err != nil {
		return 0, err
	}

	return id, s.db.Get(&id, queryCreate, u.Handle, hash)
}

func (s *service) Delete(handle string) error {
	_, err := s.db.Exec(queryDelete, handle)
	return err
}

func (s *service) GetAuth(handle string) (people.AuthUser, error) {
	var u people.AuthUser
	return u, s.db.Get(&u, queryGetAuth, handle)
}

func (s *service) Get(handle string, id *uint) (people.User, error) {
	if id == nil {
		id = new(uint)
	}
	var u people.User
	return u, s.db.Get(&u, queryGet, id, handle)
}
