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
	MessageFields = "message_id, content, from_id, thread_id, sent_at"
	SelectMessage = "SELECT " + MessageFields + " FROM message"
)

func (r *messageRepo) Create(threadID uint, content string, fromID uint) (people.DBMessage, error) {
	const query = "INSERT INTO message(thread_id, content, from_id) VALUES ($1, $2, $3) RETURNING " + MessageFields
	var m people.DBMessage
	if err := r.db.Get(&m, query, threadID, content, fromID); err != nil {
		return m, fmt.Errorf("Message.Create: %w", err)
	}
	return m, nil
}

func (r *messageRepo) ListThreadMessages(threadID uint, p pagination.ID) ([]people.DBMessage, error) {
	q, args, err := NewQuery(SelectMessage).
		Where("thread_id = ?", threadID).
		Paginate(p, "message_id", "?").
		Build()
	if err != nil {
		return nil, fmt.Errorf("Message.ListThreadMessages: %w", err)
	}
	var msgs []people.DBMessage
	if err := r.db.Select(&msgs, q, args...); err != nil {
		return nil, fmt.Errorf("Message.ListThreadMessages: %w", err)
	}
	return msgs, nil
}

func (r *messageRepo) CreateThread(userIDs ...uint) (uint, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Message.CreateThread: %w", err)
	}
	defer tx.Rollback()
	var threadID uint
	if err := r.db.Get(&threadID, "INSERT INTO thread DEFAULT VALUES RETURNING thread_id"); err != nil {
		return 0, fmt.Errorf("Message.CreateThread: %w", err)
	}
	values := make([]string, len(userIDs))
	for i, userID := range userIDs {
		values[i] = fmt.Sprintf("(%v, %v)", threadID, userID)
	}
	q, args, err := NewQuery("INSERT INTO thread_user(thread_id, user_id)").Values(values...).Build()
	if err != nil {
		return 0, fmt.Errorf("Message.CreateThread: %w", err)
	}
	if _, err := r.db.Exec(q, args...); err != nil {
		return 0, fmt.Errorf("Message.CreateThread: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("Message.CreateThread: %w", err)
	}
	return threadID, nil
}

func (r *messageRepo) GetThreadID(userIDs ...uint) (uint, error) {
	q, args, err := NewQuery("SELECT thread_id FROM thread_user tu").
		Where("user_id IN (?)", userIDs).
		GroupBy("thread_id").Having("COUNT(user_id) = (SELECT COUNT(user_id) FROM thread_user WHERE thread_id = tu.thread_id)").
		Build()
	if err != nil {
		return 0, fmt.Errorf("Message.GetThreadID: %w", err)
	}
	var id uint
	if err := r.db.Get(&id, q, args...); err != nil {
		return 0, fmt.Errorf("Message.GetThreadID: %w", err)
	}
	return id, nil
}
func (r *messageRepo) GetThreadUsers(threadID uint) ([]uint, error) {
	const query = "SELECT user_id FROM thread_user WHERE thread_id = $1"
	var ids []uint
	err := r.db.Select(&ids, query, threadID)
	if err != nil {
		return nil, fmt.Errorf("Message.GetThreadUsers: %w", err)
	}
	return ids, nil
}

func (r *messageRepo) GetLatestMessage(threadID uint) (people.DBMessage, error) {
	q, args, err := NewQuery(SelectMessage).
		Where("thread_id = ?", threadID).
		OrderBy(false, "message_id").
		Limit(1).
		Build()
	if err != nil {
		return people.DBMessage{}, fmt.Errorf("Message.GetLatestMessage: %w", err)
	}
	var msg people.DBMessage
	if err := r.db.Get(&msg, q, args...); err != nil {
		return people.DBMessage{}, fmt.Errorf("Message.GetLatestMessage: %w", err)
	}
	return msg, nil
}

func (r *messageRepo) ListThreadIDs(userID uint, p pagination.ID) ([]uint, error) {
	q, args, err := NewQuery("SELECT message.thread_id FROM message").
		Join("thread_user", "message.thread_id = thread_user.thread_id").
		Where("thread_user.user_id = ?", userID).
		GroupBy("message.thread_id").
		Paginate(p, "MAX(message_id)", "(SELECT MAX(message_id) FROM message WHERE thread_id = ?)").
		Build()
	if err != nil {
		return nil, fmt.Errorf("Message.ListThreads: %w", err)
	}
	ids := make([]uint, p.Limit)
	if err := r.db.Select(&ids, q, args...); err != nil {
		return nil, fmt.Errorf("Message.ListThreads: %w", err)
	}
	return ids, nil
}

func (r *messageRepo) ListLatestMessages(threadIDs ...uint) ([]people.DBMessage, error) {
	if len(threadIDs) == 0 {
		return []people.DBMessage{}, nil
	}
	q, args, err := NewQuery(SelectMessage).
		Where("thread_id IN (?)", threadIDs).
		Build()
	if err != nil {
		return nil, fmt.Errorf("Message.ListLatestMessages: %w", err)
	}
	msgs := make([]people.DBMessage, len(threadIDs))
	if err := r.db.Select(&msgs, q, args...); err != nil {
		return nil, fmt.Errorf("Message.ListLatestMessages: %w", err)
	}
	return msgs, nil
}

func (r *messageRepo) ListThreadUsers(threadIDs ...uint) ([]people.ThreadUser, error) {
	if len(threadIDs) == 0 {
		return []people.ThreadUser{}, nil
	}
	q, args, err := NewQuery("SELECT thread_id, user_id FROM thread_user").
		Where("thread_id IN (?)", threadIDs).
		Build()
	if err != nil {
		return nil, fmt.Errorf("Message.ListThreadUsers: %w", err)
	}
	var users []people.ThreadUser
	if err := r.db.Select(&users, q, args...); err != nil {
		return nil, fmt.Errorf("Message.ListThreadUsers: %w", err)
	}
	return users, nil
}
