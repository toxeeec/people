package inmem

import (
	"time"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type messageRepo struct {
	m      map[uint]people.DBMessage
	lastID uint
}

func (r *messageRepo) newID() uint {
	r.lastID++
	return r.lastID
}

func NewMessageRepository(m map[uint]people.DBMessage) repository.Message {
	return &messageRepo{m: m}
}

func (r *messageRepo) Create(message people.Message, from uint, to uint) (people.DBMessage, error) {
	id := r.newID()
	r.m[id] = people.DBMessage{Message: message, From: from, To: to, SentAt: time.Now()}
	return r.m[id], nil
}
