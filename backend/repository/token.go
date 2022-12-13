package repository

import (
	"github.com/google/uuid"
	people "github.com/toxeeec/people/backend"
)

type Token interface {
	Create(people.RefreshToken) error
	Get(value string) (people.RefreshToken, error)
	Delete(uuid.UUID) error
	Update(people.RefreshToken) error
}
