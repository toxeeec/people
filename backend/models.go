package people

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PaginationMeta[T any] struct {
	Oldest T `json:"oldest"`
	Newest T `json:"newest"`
}

type PaginatedResults[T Identifier[U], U any] struct {
	Data []T                `json:"data"`
	Meta *PaginationMeta[U] `json:"meta,omitempty"`
}

type ErrorKind uint

//go:generate stringer -type=ErrorKind
const (
	ValidationError ErrorKind = iota
	AuthError
	NotFoundError
	ConflictError
	ResourceError
	InternalError
)

func (e *Error) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

func (e *ErrorKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

type RefreshToken struct {
	ID     uuid.UUID `db:"token_id"`
	Value  string    `db:"value"`
	UserID uint      `db:"user_id"`
}

type Image struct {
	ID        uint      `db:"image_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UserID    uint      `db:"user_id"`
	InUse     bool      `db:"in_use"`
}

type UserMessage struct {
	Content  string `fake:"{sentence}" json:"content"`
	ThreadID uint   `fake:"skip" json:"threadID"`
}

type DBMessage struct {
	ID       uint      `db:"message_id" fake:"skip"`
	Content  string    `db:"content" fake:"{sentence}"`
	FromID   uint      `db:"from_id" fake:"skip"`
	ThreadID uint      `db:"thread_id" fake:"skip"`
	SentAt   time.Time `db:"sent_at" fake:"skip"`
}

type NotificationType string

const (
	MessageNotification = "message"
)

type Notification struct {
	Type NotificationType `json:"type"`
	Data any              `json:"data"`
	To   uint             `json:"-"`
}

type ThreadUser struct {
	ID     uint `db:"thread_id"`
	UserID uint `db:"user_id"`
}
