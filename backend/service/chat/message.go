package chat

import (
	"strings"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service"
)

func trim(m people.UserMessage) people.UserMessage {
	m.Message.Message = strings.TrimSpace(m.Message.Message)
	return m
}

func validate(m people.UserMessage) error {
	if len(m.Message.Message) == 0 {
		return service.NewError(people.ValidationError, "Message cannot be empty")
	}
	return nil
}
