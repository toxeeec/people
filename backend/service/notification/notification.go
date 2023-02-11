package notification

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type Service interface {
	Notify(notifType people.NotificationType, from, to uint, content *people.ServerMessage) error
}

type notificationService struct {
	channel chan<- people.Notification
	ur      repository.User
}

func NewService(channel chan<- people.Notification, ur repository.User) Service {
	return &notificationService{channel, ur}
}

func (s *notificationService) Notify(notifType people.NotificationType, from, to uint, content *people.ServerMessage) error {
	notif := people.Notification{
		Type:    notifType,
		From:    from,
		To:      to,
		Content: content,
	}
	s.channel <- notif
	return nil
}
