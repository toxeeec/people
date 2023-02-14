package repotest

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type MessageSuite struct {
	suite.Suite
	repo repository.Message
	ur   repository.User
	fns  TestFns
}

func NewMessageSuite(mr repository.Message, ur repository.User, fns TestFns) *MessageSuite {
	return &MessageSuite{repo: mr, ur: ur, fns: fns}
}

func (s *MessageSuite) TestCreate() {
	var au people.AuthUser
	var m people.Message
	gofakeit.Struct(&au)
	gofakeit.Struct(&m)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)

	dbm, err := s.repo.Create(m, u1.ID, u2.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), m.Content, dbm.Content)
	assert.Equal(s.T(), u1.ID, dbm.From)
	assert.Equal(s.T(), u2.ID, dbm.To)

	dbm, err = s.repo.Create(m, u1.ID, u1.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), m.Content, dbm.Content)
	assert.Equal(s.T(), u1.ID, dbm.From)
	assert.Equal(s.T(), u1.ID, dbm.To)
}

func (s *MessageSuite) TestListUserMessages() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u3, _ := s.ur.Create(au)

	userMessages := 5
	var m people.Message
	for i := 0; i < userMessages; i++ {
		gofakeit.Struct(&m)
		s.repo.Create(m, u1.ID, u2.ID)
		gofakeit.Struct(&m)
		s.repo.Create(m, u2.ID, u1.ID)
		gofakeit.Struct(&m)
		s.repo.Create(m, u3.ID, u1.ID)

		gofakeit.Struct(&m)
		s.repo.Create(m, u3.ID, u3.ID)
	}

	dbms, err := s.repo.ListUserMessages(u1.ID, u2.ID, pagination.ID{Limit: 20})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), dbms, userMessages*2)

	dbms, err = s.repo.ListUserMessages(u3.ID, u3.ID, pagination.ID{Limit: 20})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), dbms, userMessages)
}

func (s *MessageSuite) SetupTest() {
	if s.fns.SetupTest != nil {
		s.fns.SetupTest()
	}
}
