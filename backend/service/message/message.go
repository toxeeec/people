package message

import (
	"context"
	"encoding/json"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service"
	"github.com/toxeeec/people/backend/service/notification"
	"github.com/toxeeec/people/backend/service/user"
)

type Service interface {
	ReadMessage(from uint, data []byte) error
	ListUserMessages(handle string, userID uint, params pagination.IDParams) (people.UserMessages, error)
}

type messageService struct {
	mr repository.Message
	ur repository.User
	ns notification.Service
	us user.Service
}

func NewService(mr repository.Message, ur repository.User, ns notification.Service, us user.Service) Service {
	return &messageService{mr, ur, ns, us}
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
	sm := people.IntoServerMessage(dbm, from.Handle, um.To)
	return s.ns.Notify(people.MessageNotification, from.ID, to, &sm)
}

func (s *messageService) ListUserMessages(handle string, userID uint, params pagination.IDParams) (people.UserMessages, error) {
	p := pagination.New(params.Before, params.After, params.Limit)
	target, err := s.us.GetUser(context.Background(), handle, userID, true)
	if err != nil {
		return people.UserMessages{}, err
	}
	u, err := s.ur.Get(userID)
	if err != nil {
		return people.UserMessages{}, err
	}
	dbms, err := s.mr.ListUserMessages(userID, target.ID, p)
	if err != nil {
		return people.UserMessages{}, err
	}
	sms := people.IntoServerMessages(dbms, map[uint]string{target.ID: target.Handle, u.ID: u.Handle})
	return people.UserMessages{User: target, Data: pagination.NewResults[people.ServerMessage, uint](sms)}, nil
}
