package inmem

import (
	"math"
	"time"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
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

func (r *messageRepo) ListUserMessages(user1ID, user2ID uint, p pagination.ID) ([]people.DBMessage, error) {
	before := uint(math.MaxUint)
	if p.Before != nil {
		before = *p.Before
	}
	after := uint(0)
	if p.After != nil {
		after = *p.After
	}
	var ms []people.DBMessage
	for k, v := range r.m {
		if ((v.From == user1ID && v.To == user2ID) || (v.From == user2ID && v.To == user1ID)) && k < before && k > after {
			ms = append(ms, v)
		}
		if len(ms) == int(p.Limit) {
			break
		}
	}
	return ms, nil
}
