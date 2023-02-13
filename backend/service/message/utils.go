package message

import (
	"strings"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service"
)

func trim(m people.UserMessage) people.UserMessage {
	m.Content = strings.TrimSpace(m.Content)
	return m
}

func validate(m people.UserMessage) error {
	if len(m.Content) == 0 {
		return service.NewError(people.ValidationError, "Content cannot be empty")
	}
	return nil
}
