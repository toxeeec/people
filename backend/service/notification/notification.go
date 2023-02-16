package notification

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type Service interface {
	Notify(notifType people.NotificationType, data any, ids ...uint)
}

type notificationService struct {
	channel chan<- people.Notification
	ur      repository.User
}

func NewService(channel chan<- people.Notification, ur repository.User) Service {
	return &notificationService{channel, ur}
}

func (s *notificationService) Notify(notifType people.NotificationType, data any, ids ...uint) {
	for _, id := range ids {
		notif := people.Notification{
			Type: notifType,
			Data: data,
			To:   id,
		}
		s.channel <- notif
	}
}
