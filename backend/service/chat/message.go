package chat

import (
	"strings"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service"
)

func trim(m people.Message) people.Message {
	m.Message = strings.TrimSpace(m.Message)
	return m
}

func validate(m people.Message) error {
	if len(m.Message) == 0 {
		return service.NewError(people.ValidationError, "Message cannot be empty")
	}
	return nil
}
