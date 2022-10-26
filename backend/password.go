package people

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password string

var ErrInvalidPassword = errors.New("invalid password")

func (p Password) Hash() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(bytes), err
}

func (p Password) Compare(hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	if err != nil {
		return ErrInvalidPassword
	}

	return nil
}
