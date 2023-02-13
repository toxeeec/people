package repotest

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
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
	var au1 people.AuthUser
	var au2 people.AuthUser
	var m people.Message
	gofakeit.Struct(&au1)
	gofakeit.Struct(&au2)
	gofakeit.Struct(&m)
	u1, _ := s.ur.Create(au1)
	u2, _ := s.ur.Create(au2)

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

func (s *MessageSuite) SetupTest() {
	if s.fns.SetupTest != nil {
		s.fns.SetupTest()
	}
}
