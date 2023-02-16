package repository

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
)

type Message interface {
	Create(threadID uint, content string, fromID uint) (people.DBMessage, error)
	CreateThread(userIDs ...uint) (uint, error)
	GetThreadID(userIDs ...uint) (uint, error)
	ListThreadIDs(userID uint, p pagination.ID) ([]uint, error)
	GetThreadUsers(threadID uint) ([]uint, error)
	ListThreadUsers(threadIDs ...uint) ([]people.ThreadUser, error)
	GetLatestMessage(threadID uint) (people.DBMessage, error)
	ListLatestMessages(threadIDs ...uint) ([]people.DBMessage, error)
	ListThreadMessages(threadID uint, p pagination.ID) ([]people.DBMessage, error)
}
