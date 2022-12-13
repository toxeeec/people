package repository

import people "github.com/toxeeec/people/backend"

type User interface {
	GetID(handle string) (uint, error)
	Create(people.AuthUser) (people.User, error)
	Delete(id uint) error
	GetHash(id uint) (string, error)
	Get(id uint) (people.User, error)
	List(ids []uint) ([]people.User, error)
}
