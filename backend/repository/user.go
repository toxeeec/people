package repository

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
)

type User interface {
	GetID(handle string) (uint, error)
	Create(people.AuthUser) (people.User, error)
	Delete(id uint) error
	GetHash(id uint) (string, error)
	Get(id uint) (people.User, error)
	List(ids []uint) ([]people.User, error)
	ListIDs(handles ...string) ([]uint, error)
	ListMatches(query string, p pagination.ID) ([]people.User, error)
	Update(userID uint, handle string) (people.User, error)
}
