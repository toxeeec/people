package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type messageRepo struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) repository.Message {
	return &messageRepo{db}
}

const (
	MessageFields = "message_id, from_id, to_id, content, sent_at"
	SelectMessage = "SELECT " + MessageFields + " FROM message"
)

func (r *messageRepo) Create(message people.Message, from uint, to uint) (people.DBMessage, error) {
	const query = "INSERT INTO message(from_id, to_id, content) VALUES ($1, $2, $3) RETURNING " + MessageFields
	var m people.DBMessage
	if err := r.db.Get(&m, query, from, to, message.Content); err != nil {
		return m, fmt.Errorf("Message.Create: %w", err)
	}
	return m, nil
}

func (r *messageRepo) ListUserMessages(user1ID, user2ID uint, p pagination.ID) ([]people.DBMessage, error) {
	q, args, err := NewQuery(SelectMessage).
		Where("((from_id = ? AND to_id = ?) OR (to_id = ? AND from_id = ?))", user1ID, user2ID, user1ID, user2ID).
		Paginate(p, "message_id", "?").
		Build()
	if err != nil {
		return nil, fmt.Errorf("Message.ListUserMessages: %w", err)
	}
	ms := make([]people.DBMessage, p.Limit)
	if err := r.db.Select(&ms, q, args...); err != nil {
		return nil, fmt.Errorf("Message.ListUserMessages: %w", err)
	}
	return ms, nil
}
