package repository

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
)

type Message interface {
	Create(message people.Message, from uint, to uint) (people.DBMessage, error)
	ListUserMessages(user1ID, user2ID uint, p pagination.ID) ([]people.DBMessage, error)
}
