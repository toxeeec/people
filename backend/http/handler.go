package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/toxeeec/people/backend/service/auth"
	"github.com/toxeeec/people/backend/service/image"
	"github.com/toxeeec/people/backend/service/post"
	"github.com/toxeeec/people/backend/service/user"
)

type handler struct {
	as auth.Service
	us user.Service
	ps post.Service
	is image.Service
}

func NewHandler(v *validator.Validate, as auth.Service, us user.Service, ps post.Service, is image.Service) *handler {
	var h handler
	h.as = as
	h.us = us
	h.ps = ps
	h.is = is
	return &h
}
