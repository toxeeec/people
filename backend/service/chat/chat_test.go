package chat_test

import (
	"encoding/json"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/service/chat"
	"github.com/toxeeec/people/backend/service/notification"
)

type ChatSuite struct {
	suite.Suite
	cs chat.Service
	ur repository.User
}

func (s *ChatSuite) TestReadMessage() {
	var au1 people.AuthUser
	var au2 people.AuthUser
	gofakeit.Struct(&au1)
	gofakeit.Struct(&au2)
	u1, _ := s.ur.Create(au1)
	u2, _ := s.ur.Create(au2)

	validationError := people.ValidationError

	tests := map[string]struct {
		message people.Message
		valid   bool
		kind    *people.ErrorKind
	}{
		"empty message": {people.Message{Message: "", To: u2.Handle}, false, &validationError},
		"valid":         {people.Message{Message: gofakeit.SentenceSimple(), To: u2.Handle}, true, nil},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			msg, _ := json.Marshal(tc.message)
			err := s.cs.ReadMessage(u1.ID, msg)
			assert.Equal(s.T(), tc.valid, err == nil)
			if !tc.valid {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}

		})
	}
}

func (s *ChatSuite) SetupTest() {
	um := map[uint]people.User{}
	ns := notification.NewService()
	s.ur = inmem.NewUserRepository(um)
	s.cs = chat.NewService(s.ur, ns)

}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(ChatSuite))
}
