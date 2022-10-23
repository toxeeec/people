package people

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var v = validator.New()

var (
	ErrUnknown        = errors.New("Unknown error")
	ErrHandleAlphanum = errors.New("Handle can only contain letters and numbers")
)

func (u AuthUser) Validate() error {
	if err := v.Var(u.Handle, "alphanum"); err != nil {
		err := err.(validator.ValidationErrors)
		switch err[0].Tag() {
		case "alphanum":
			return ErrHandleAlphanum
		default:
			return ErrUnknown
		}
	}
	return nil
}
