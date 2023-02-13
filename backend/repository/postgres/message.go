package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
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
