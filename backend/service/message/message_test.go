package message_test

import (
	"context"
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
	"github.com/toxeeec/people/backend/service/image"
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
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	t1, _ := s.mr.CreateThread(u1.ID)

	validationError := people.ValidationError

	tests := map[string]struct {
		message people.UserMessage
		valid   bool
		kind    *people.ErrorKind
	}{
		"empty message":  {people.UserMessage{Content: "", ThreadID: t1}, false, &validationError},
		"invalid thread": {people.UserMessage{Content: gofakeit.SentenceSimple(), ThreadID: 0}, false, &validationError},
		"valid":          {people.UserMessage{Content: gofakeit.SentenceSimple(), ThreadID: t1}, true, nil},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			msg, _ := json.Marshal(tc.message)
			err := s.ms.ReadMessage(context.Background(), u1.ID, msg)
			assert.Equal(s.T(), tc.valid, err == nil)
			if !tc.valid {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}

		})
	}
}

func (s *MessageSuite) TestGetThread() {
	var unknownUser people.AuthUser
	var au people.AuthUser
	gofakeit.Struct(&unknownUser)
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)

	t1, _ := s.mr.CreateThread(u1.ID)
	t2, _ := s.mr.CreateThread(u2.ID)

	authError := people.AuthError
	notFoundError := people.NotFoundError

	tests := map[string]struct {
		threadID uint
		valid    bool
		kind     *people.ErrorKind
	}{
		"no permission": {t2, false, &authError},
		"valid":         {t1, true, &notFoundError},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			t, err := s.ms.GetThread(context.Background(), u1.ID, tc.threadID)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(s.T(), tc.threadID, t.ID)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}

		})
	}
}

func (s *MessageSuite) TestGetUsersThread() {
	var unknownUser people.AuthUser
	var au people.AuthUser
	gofakeit.Struct(&unknownUser)
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)

	var m people.Message
	t1, _ := s.mr.CreateThread(u1.ID, u2.ID)
	t2, _ := s.mr.CreateThread(u1.ID)
	gofakeit.Struct(&m)
	s.mr.Create(t1, m.Content, u1.ID)
	gofakeit.Struct(&m)
	msg, _ := s.mr.Create(t1, m.Content, u2.ID)
	t1Latest := message.IntoMessage(msg, u2)
	gofakeit.Struct(&m)
	s.mr.Create(t2, m.Content, u1.ID)
	gofakeit.Struct(&m)
	msg, _ = s.mr.Create(t2, m.Content, u1.ID)
	t2Latest := message.IntoMessage(msg, u1)

	notFoundError := people.NotFoundError

	tests := map[string]struct {
		handle string
		latest *people.Message
		valid  bool
		kind   *people.ErrorKind
	}{
		"unknown handle": {unknownUser.Handle, nil, false, &notFoundError},
		"valid":          {u2.Handle, &t1Latest, true, &notFoundError},
		"valid(self)":    {u1.Handle, &t2Latest, true, &notFoundError},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			t, err := s.ms.GetUsersThread(context.Background(), u1.ID, tc.handle)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(s.T(), tc.latest.ThreadID, t.ID)
				assert.Equal(s.T(), t.Latest.Content, tc.latest.Content)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}

		})
	}
}

func (s *MessageSuite) TestListThreadMessages() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)
	t1, _ := s.mr.CreateThread(u1.ID, u2.ID)
	t2, _ := s.mr.CreateThread(u1.ID)

	var m people.Message
	var messages [6]people.DBMessage
	for i := 0; i < len(messages)/2; i++ {
		gofakeit.Struct(&m)
		messages[i*2], _ = s.mr.Create(t1, m.Content, u1.ID)
		gofakeit.Struct(&m)
		messages[i*2+1], _ = s.mr.Create(t1, m.Content, u2.ID)
	}
	gofakeit.Struct(&m)
	s.mr.Create(t2, m.Content, u1.ID)

	limit := uint(10)
	res, err := s.ms.ListThreadMessages(context.Background(), t1, u1.ID, pagination.IDParams{Limit: &limit})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), res.Data, len(messages))
}

func (s *MessageSuite) TestListThreads() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u3, _ := s.ur.Create(au)
	var threads [3]uint
	threads[0], _ = s.mr.CreateThread(u1.ID)
	threads[1], _ = s.mr.CreateThread(u1.ID, u2.ID)
	threads[2], _ = s.mr.CreateThread(u1.ID, u2.ID, u3.ID)
	var m people.Message
	for _, t := range threads {
		gofakeit.Struct(&m)
		s.mr.Create(t, m.Content, u1.ID)
	}

	limit := uint(10)
	actual, err := s.ms.ListThreads(context.Background(), u1.ID, pagination.IDParams{Limit: &limit})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), actual.Data, len(threads))
	for _, t := range actual.Data {
		assert.Contains(s.T(), threads, t.ID)
	}
}

func (s *MessageSuite) SetupTest() {
	um := map[uint]people.User{}
	fm := map[inmem.FollowKey]time.Time{}
	lm := map[inmem.LikeKey]struct{}{}
	pm := map[uint]people.Post{}
	im := map[uint]people.Image{}
	msgs := make(map[uint][]people.DBMessage)
	threads := make(map[uint]struct{})
	threadUsers := make(map[uint][]uint)
	v := validator.New()
	s.mr = inmem.NewMessageRepository(msgs, threads, threadUsers, um)
	s.ur = inmem.NewUserRepository(um)
	fr := inmem.NewFollowRepository(fm, um)
	lr := inmem.NewLikeRepository(lm, pm, um)
	ns := notification.NewService(make(chan people.Notification, 32), s.ur)
	ir := inmem.NewImageRepository(im)
	is := image.NewService(ir)
	us := user.NewService(v, s.ur, fr, lr, is)
	s.ms = message.NewService(s.mr, s.ur, ns, us)
}

func TestMessageSuite(t *testing.T) {
	suite.Run(t, new(MessageSuite))
}
