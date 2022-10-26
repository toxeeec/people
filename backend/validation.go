package people

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var v = validator.New()

var (
	ErrUnknown            = errors.New("Unknown error")
	ErrSpecialCharsHandle = errors.New("Handle cannot contain special characters")
	ErrEmptyContent       = errors.New("Content cannot be empty")
)

func (u AuthUser) Validate() error {
	if err := v.Var(u.Handle, "alphanum"); err != nil {
		err := err.(validator.ValidationErrors)
		switch err[0].Tag() {
		case "alphanum":
			return ErrSpecialCharsHandle
		default:
			return ErrUnknown
		}
	}
	return nil
}

func (p PostBody) Validate() error {
	if err := v.Var(p.Content, "min=1"); err != nil {
		err := err.(validator.ValidationErrors)
		switch err[0].Tag() {
		case "min":
			return ErrEmptyContent
		default:
			return ErrUnknown
		}
	}
	return nil
}
