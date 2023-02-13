package message

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

type messageService struct {
	mr repository.Message
	ur repository.User
	ns notification.Service
}

func NewService(mr repository.Message, ur repository.User, ns notification.Service) Service {
	return &messageService{mr, ur, ns}
}

func (s *messageService) ReadMessage(fromID uint, data []byte) error {
	var um people.UserMessage
	if err := json.Unmarshal(data, &um); err != nil {
		return err
	}
	um = trim(um)
	if err := validate(um); err != nil {
		return err
	}
	to, err := s.ur.GetID(um.To)
	if err != nil {
		return service.NewError(people.NotFoundError, "User not found")
	}
	from, err := s.ur.Get(fromID)
	if err != nil {
		return err
	}
	msg := people.Message{Content: um.Content}
	dbm, err := s.mr.Create(msg, fromID, to)
	if err != nil {
		println(err.Error())
		return service.InternalServerError
	}
	sm := &people.ServerMessage{Message: dbm.Message, From: from.Handle, To: um.To}
	return s.ns.Notify(people.MessageNotification, from.ID, to, sm)
}
