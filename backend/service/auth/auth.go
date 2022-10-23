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
