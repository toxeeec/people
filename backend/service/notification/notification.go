package notification

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service"
)

type Service interface {
	Notify(notifType people.NotificationType, from, to uint, content *people.UserMessage) error
}

type notificationService struct {
	channel chan<- people.Notification
	ur      repository.User
}

func NewService(channel chan<- people.Notification, ur repository.User) Service {
	return &notificationService{channel, ur}
}

func (s *notificationService) Notify(notifType people.NotificationType, from, to uint, content *people.UserMessage) error {
	u, err := s.ur.Get(from)
	if err != nil {
		return service.InternalServerError
	}
	notif := people.Notification{
		Type: notifType,
		From: u.Handle,
		To:   to,
	}
	if content != nil {
		notif.Content = &content.Message
	}
	s.channel <- notif
	return nil
}
