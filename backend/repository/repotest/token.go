package repotest

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service/auth"
)

type TokenSuite struct {
	suite.Suite
	repo repository.Token
	ur   repository.User
	fns  TestFns
}

func NewTokenSuite(tr repository.Token, ur repository.User, fns TestFns) *TokenSuite {
	return &TokenSuite{repo: tr, ur: ur, fns: fns}
}

func (s *TokenSuite) TestCreate() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	rt, _ := auth.NewRefreshToken(u.ID, nil)

	err := s.repo.Create(rt)
	assert.NoError(s.T(), err)
	actual, _ := s.repo.Get(rt.Value)
	assert.Equal(s.T(), rt, actual)
}

func (s *TokenSuite) TestGet() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	rt, _ := auth.NewRefreshToken(u.ID, nil)
	s.repo.Create(rt)

	actual, err := s.repo.Get(rt.Value)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), rt, actual)
}

func (s *TokenSuite) TestDelete() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	rt, _ := auth.NewRefreshToken(u.ID, nil)
	s.repo.Create(rt)

	err := s.repo.Delete(rt.ID)
	assert.NoError(s.T(), err)
	_, err = s.repo.Get(rt.Value)
	assert.Error(s.T(), err)
}

func (s *TokenSuite) TestUpdate() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	rt, _ := auth.NewRefreshToken(u.ID, nil)
	s.repo.Create(rt)

	time.Sleep(time.Second)
	newRT, _ := auth.NewRefreshToken(u.ID, &rt.ID)
	err := s.repo.Update(newRT)
	assert.NoError(s.T(), err)
	_, err = s.repo.Get(rt.Value)
	assert.Error(s.T(), err)
	actual, err := s.repo.Get(newRT.Value)
	assert.Equal(s.T(), newRT, actual)
	assert.NoError(s.T(), err)
}

func (s *TokenSuite) SetupTest() {
	if s.fns.SetupTest != nil {
		s.fns.SetupTest()
	}
}
