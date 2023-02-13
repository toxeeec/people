package repository

import people "github.com/toxeeec/people/backend"

type Message interface {
	Create(message people.Message, from uint, to uint) (people.DBMessage, error)
}
