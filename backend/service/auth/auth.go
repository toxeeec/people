package auth

import (
	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
)

type service struct {
	db *sqlx.DB
	us people.UserService
}

func NewService(db *sqlx.DB, us people.UserService) people.AuthService {
	return &service{db, us}
}

// VerifyCredentials returns id of the user.
func (s *service) VerifyCredentials(u people.AuthUser) (uint, error) {
	expected, err := s.us.Get(u.Handle)
	if err != nil {
		return 0, err
	}

	if err := u.Password.Compare(*expected.Hash); err != nil {
		return 0, err
	}

	return *expected.ID, nil
}
