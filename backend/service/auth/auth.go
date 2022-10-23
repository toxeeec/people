package auth

import (
	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) people.AuthService {
	return &service{db}
}
