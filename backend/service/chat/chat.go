package chat

import (
	"encoding/json"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service"
	"github.com/toxeeec/people/backend/service/notification"
)

type Service interface {
	ReadMessage(from uint, data []byte) error
}

type chatService struct {
	ur repository.User
	ns notification.Service
}

func NewService(ur repository.User, ns notification.Service) Service {
	return &chatService{
		ur,
		ns,
	}
}

func (s *chatService) ReadMessage(from uint, data []byte) error {
	var msg people.UserMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}
	msg = trim(msg)
	if err := validate(msg); err != nil {
		return err
	}
	to, err := s.ur.GetID(msg.To)
	if err != nil {
		return service.NewError(people.NotFoundError, "User not found")
	}
	// TODO: save message
	return s.ns.Notify(people.MessageNotification, from, to, &msg)
}
