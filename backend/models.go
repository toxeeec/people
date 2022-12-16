package people

import (
	"encoding/json"
	"fmt"

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