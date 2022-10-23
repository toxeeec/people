package handler

import (
	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/auth"
	"github.com/toxeeec/people/backend/service/user"
)

type handler struct {
	as people.AuthService
	us people.UserService
}

func New(db *sqlx.DB) people.ServerInterface {
	var h handler
	h.us = user.NewService(db)
	h.as = auth.NewService(db, h.us)
	return &h
}
