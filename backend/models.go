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
	Message
	To string `json:"to"`
}

func (m ServerMessage) Identify() uint {
	return m.ID
}

type DBMessage struct {
	Message
	ID     uint      `db:"message_id" fake:"skip"`
	From   uint      `db:"from_id" fake:"skip"`
	To     uint      `db:"to_id" fake:"skip"`
	SentAt time.Time `db:"sent_at" fake:"skip"`
}

func IntoServerMessage(m DBMessage, from string, to string) ServerMessage {
	return ServerMessage{Message: m.Message, ID: m.ID, From: from, To: to, SentAt: m.SentAt}
}

func IntoServerMessages(ms []DBMessage, handles map[uint]string) []ServerMessage {
	sms := make([]ServerMessage, len(ms))
	for i, m := range ms {
		sms[i] = ServerMessage{Message: m.Message, ID: m.ID, From: handles[m.From], To: handles[m.To], SentAt: m.SentAt}
	}
	return sms
}

type MessagesResponse struct {
	PaginatedResults[ServerMessage, uint]
	User User
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
