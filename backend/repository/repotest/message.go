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
	t1, _ := s.repo.CreateThread(u1.ID, u2.ID)

	dbm, err := s.repo.Create(t1, m.Content, u1.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), t1, dbm.ThreadID)
	assert.Equal(s.T(), m.Content, dbm.Content)
	assert.Equal(s.T(), u1.ID, dbm.FromID)

	gofakeit.Struct(&m)
	t2, _ := s.repo.CreateThread(u1.ID)
	dbm, err = s.repo.Create(t2, m.Content, u1.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), t2, dbm.ThreadID)
	assert.Equal(s.T(), m.Content, dbm.Content)
	assert.Equal(s.T(), u1.ID, dbm.FromID)
}

func (s *MessageSuite) TestCreateThread() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)

	t, err := s.repo.CreateThread(u1.ID, u2.ID)
	assert.NoError(s.T(), err)
	actual, _ := s.repo.GetThreadID(u1.ID, u2.ID)
	assert.Equal(s.T(), t, actual)
}

func (s *MessageSuite) TestGetThreadID() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)
	t, _ := s.repo.CreateThread(u1.ID, u2.ID)

	actual, err := s.repo.GetThreadID(u1.ID, u2.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), t, actual)
}

func (s *MessageSuite) ListThreadIDs() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u3, _ := s.ur.Create(au)

	t1, _ := s.repo.CreateThread(u1.ID)
	t2, _ := s.repo.CreateThread(u1.ID, u2.ID)
	t3, _ := s.repo.CreateThread(u1.ID, u2.ID, u3.ID)

	actual, err := s.repo.ListThreadIDs(u1.ID, pagination.ID{Limit: 10})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), actual, 3)
	assert.Contains(s.T(), actual, t1)
	assert.Contains(s.T(), actual, t2)
	assert.Contains(s.T(), actual, t3)
}

func (s *MessageSuite) TestGetThreadUsers() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u3, _ := s.ur.Create(au)
	t, _ := s.repo.CreateThread(u1.ID, u2.ID, u3.ID)

	users, err := s.repo.GetThreadUsers(t)
	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, 3)
}

func (s *MessageSuite) TestListThreadUsers() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u3, _ := s.ur.Create(au)
	t1, _ := s.repo.CreateThread(u1.ID)
	t2, _ := s.repo.CreateThread(u2.ID, u3.ID)
	t3, _ := s.repo.CreateThread(u1.ID, u2.ID, u3.ID)

	users, err := s.repo.ListThreadUsers(t1, t2, t3)
	assert.NoError(s.T(), err)
	var ids int
	idsMap := make(map[uint]struct{})
	for _, user := range users {
		if _, ok := idsMap[user.UserID]; !ok {
			ids++
			idsMap[user.UserID] = struct{}{}
		}
	}
	assert.Equal(s.T(), ids, 3)
}

func (s *MessageSuite) TestGetLatestMessage() {
	var au people.AuthUser
	var m people.Message
	gofakeit.Struct(&au)
	gofakeit.Struct(&m)
	u1, _ := s.ur.Create(au)
	t, _ := s.repo.CreateThread(u1.ID)
	gofakeit.Struct(&m)
	s.repo.Create(t, m.Content, u1.ID)
	dbm, _ := s.repo.Create(t, m.Content, u1.ID)

	actual, err := s.repo.GetLatestMessage(t)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), dbm, actual)
}

func (s *MessageSuite) TestListLatestMessages() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u2, _ := s.ur.Create(au)
	gofakeit.Struct(&au)
	u3, _ := s.ur.Create(au)
	t1, _ := s.repo.CreateThread(u1.ID)
	t2, _ := s.repo.CreateThread(u1.ID, u2.ID)
	t3, _ := s.repo.CreateThread(u1.ID, u2.ID, u3.ID)
	var m people.Message
	gofakeit.Struct(&m)
	dbm1, _ := s.repo.Create(t1, m.Content, u1.ID)
	gofakeit.Struct(&m)
	dbm2, _ := s.repo.Create(t2, m.Content, u2.ID)
	gofakeit.Struct(&m)
	dbm3, _ := s.repo.Create(t3, m.Content, u3.ID)

	actual, err := s.repo.ListLatestMessages(t1, t2, t3)
	assert.NoError(s.T(), err)
	assert.Len(s.T(), actual, 3)
	assert.Contains(s.T(), actual, dbm1)
	assert.Contains(s.T(), actual, dbm2)
	assert.Contains(s.T(), actual, dbm3)
}

func (s *MessageSuite) TestListThreadMessages() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	t, _ := s.repo.CreateThread(u.ID)
	var m people.Message
	gofakeit.Struct(&m)
	dbm1, _ := s.repo.Create(t, m.Content, u.ID)
	gofakeit.Struct(&m)
	dbm2, _ := s.repo.Create(t, m.Content, u.ID)
	gofakeit.Struct(&m)
	dbm3, _ := s.repo.Create(t, m.Content, u.ID)

	actual, err := s.repo.ListThreadMessages(t, pagination.ID{Limit: 10})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), actual, 3)
	assert.Contains(s.T(), actual, dbm1)
	assert.Contains(s.T(), actual, dbm2)
	assert.Contains(s.T(), actual, dbm3)
}
