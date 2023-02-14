package message_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/service/message"
	"github.com/toxeeec/people/backend/service/notification"
	"github.com/toxeeec/people/backend/service/user"
)

type MessageSuite struct {
	suite.Suite
	ms message.Service
	mr repository.Message
	ur repository.User
}

func (s *MessageSuite) TestReadMessage() {
	var au1 people.AuthUser
	var au2 people.AuthUser
	gofakeit.Struct(&au1)
	gofakeit.Struct(&au2)
	u1, _ := s.ur.Create(au1)
	u2, _ := s.ur.Create(au2)

	validationError := people.ValidationError

	tests := map[string]struct {
		message people.UserMessage
		valid   bool
		kind    *people.ErrorKind
	}{
		"empty message": {people.UserMessage{Message: people.Message{Content: ""}, To: u2.Handle}, false, &validationError},
		"valid":         {people.UserMessage{Message: people.Message{Content: gofakeit.SentenceSimple()}, To: u2.Handle}, true, nil},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			msg, _ := json.Marshal(tc.message)
			err := s.ms.ReadMessage(u1.ID, msg)
			assert.Equal(s.T(), tc.valid, err == nil)
			if !tc.valid {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}

		})
	}
}

func (s *MessageSuite) TestListUserMessages() {
	var unknownUser people.AuthUser
	var au1 people.AuthUser
	var au2 people.AuthUser
	var au3 people.AuthUser
	gofakeit.Struct(&au1)
	gofakeit.Struct(&au2)
	gofakeit.Struct(&au3)
	gofakeit.Struct(&unknownUser)
	u1, _ := s.ur.Create(au1)
	u2, _ := s.ur.Create(au2)
	u3, _ := s.ur.Create(au3)
	userMessages := 5
	var m people.Message
	for i := 0; i < userMessages; i++ {
		gofakeit.Struct(&m)
		s.mr.Create(m, u1.ID, u2.ID)
		gofakeit.Struct(&m)
		s.mr.Create(m, u2.ID, u1.ID)
		gofakeit.Struct(&m)
		s.mr.Create(m, u3.ID, u1.ID)

		gofakeit.Struct(&m)
		s.mr.Create(m, u3.ID, u3.ID)
	}

	notFoundError := people.NotFoundError

	tests := map[string]struct {
		handle   string
		id       uint
		messages uint
		valid    bool
		kind     *people.ErrorKind
	}{
		"unknown handle": {unknownUser.Handle, u1.ID, 0, false, &notFoundError},
		"valid":          {u2.Handle, u1.ID, 10, true, &notFoundError},
		"valid(self)":    {u3.Handle, u3.ID, 5, true, &notFoundError},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			ums, err := s.ms.ListUserMessages(tc.handle, tc.id, pagination.IDParams{})
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(s.T(), tc.handle, ums.User.Handle)
				assert.Len(s.T(), ums.Data.Data, int(tc.messages))
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}

		})
	}
}

func (s *MessageSuite) SetupTest() {
	um := map[uint]people.User{}
	mm := map[uint]people.DBMessage{}
	fm := map[inmem.FollowKey]time.Time{}
	lm := map[inmem.LikeKey]struct{}{}
	pm := map[uint]people.Post{}
	v := validator.New()
	s.mr = inmem.NewMessageRepository(mm)
	s.ur = inmem.NewUserRepository(um)
	fr := inmem.NewFollowRepository(fm, um)
	lr := inmem.NewLikeRepository(lm, pm, um)
	ns := notification.NewService(make(chan people.Notification, 32), s.ur)
	us := user.NewService(v, s.ur, fr, lr)
	s.ms = message.NewService(s.mr, s.ur, ns, us)

}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(MessageSuite))
}
