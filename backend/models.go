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

type Message struct {
	Content string `db:"content" fake:"{sentence}" json:"content"`
}

type DBMessage struct {
	Message
	ID     uint      `db:"message_id" fake:"skip"`
	From   uint      `db:"from_id" fake:"skip"`
	To     uint      `db:"to_id" fake:"skip"`
	SentAt time.Time `db:"sent_at" fake:"skip"`
}

type UserMessage struct {
	Message
	To string `json:"to"`
}

type ServerMessage struct {
	Message
	From string `json:"from"`
	To   string `json:"to"`
}

type NotificationType string

const (
	MessageNotification = "message"
)

type Notification struct {
	Type    NotificationType `json:"type"`
	From    uint             `json:"-"`
	To      uint             `json:"-"`
	Content *ServerMessage   `json:"content,omitempty"`
}
