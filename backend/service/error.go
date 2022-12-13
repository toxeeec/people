package service

import (
	people "github.com/toxeeec/people/backend"
)

func NewError(kind people.ErrorKind, message string) error {
	return &people.Error{Kind: &kind, Message: message}
}
