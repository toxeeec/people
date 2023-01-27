package notification

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type Service interface {
	Notify(notifType people.NotificationType, from, to uint, content *people.Message)
}

type notificationService struct {
	channel chan<- people.Notification
	ur      repository.User
}

func NewService() Service {
	return &notificationService{}
}

func (s *notificationService) Notify(notifType people.NotificationType, from, to uint, content *people.Message) {
	// TODO

}
