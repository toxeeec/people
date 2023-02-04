package repotest

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type UserSuite struct {
	suite.Suite
	repo repository.User
	fns  TestFns
}

func NewUserSuite(ur repository.User, fns TestFns) *UserSuite {
	return &UserSuite{repo: ur, fns: fns}
}

func (s *UserSuite) TestGetID() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.repo.Create(au)

	id, err := s.repo.GetID(u.Handle)
	assert.Equal(s.T(), u.ID, id)
	assert.NoError(s.T(), err)
}

func (s *UserSuite) TestCreate() {
	var au people.AuthUser
	gofakeit.Struct(&au)

	u, err := s.repo.Create(au)
	assert.Equal(s.T(), u.Handle, au.Handle)
	assert.NoError(s.T(), err)
	_, err = s.repo.GetID(u.Handle)
	assert.NoError(s.T(), err)
}

func (s *UserSuite) TestDelete() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.repo.Create(au)

	err := s.repo.Delete(u.ID)
	assert.NoError(s.T(), err)
	_, err = s.repo.GetID(u.Handle)
	assert.Error(s.T(), err)
}

func (s *UserSuite) TestGetHash() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.repo.Create(au)

	h, err := s.repo.GetHash(u.ID)
	assert.Equal(s.T(), au.Password, h)
	assert.NoError(s.T(), err)
}

func (s *UserSuite) TestGet() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.repo.Create(au)

	actual, err := s.repo.Get(u.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), u, actual)
}

func (s *UserSuite) TestListMatches() {
	s.repo.Create(people.AuthUser{Handle: "abc", Password: gofakeit.Password(true, true, true, true, true, 8)})
	s.repo.Create(people.AuthUser{Handle: "DEFABCGHI", Password: gofakeit.Password(true, true, true, true, true, 8)})
	// not matching
	s.repo.Create(people.AuthUser{Handle: "defghi", Password: gofakeit.Password(true, true, true, true, true, 8)})

	ps, err := s.repo.ListMatches("abc", pagination.ID{Limit: 10})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), ps, 2)
}

func (s *UserSuite) TestUpdate() {
	var au1 people.AuthUser
	var au2 people.AuthUser
	gofakeit.Struct(&au1)
	gofakeit.Struct(&au2)
	u1, _ := s.repo.Create(au1)

	u2, err := s.repo.Update(u1.ID, au2.Handle)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), u1.ID, u2.ID)
	assert.Equal(s.T(), au2.Handle, u2.Handle)

	_, err = s.repo.GetID(u1.Handle)
	assert.Error(s.T(), err)
	u3, err := s.repo.Get(u2.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), u2.ID, u3.ID)
	assert.Equal(s.T(), u2.Handle, u3.Handle)
}

func (s *UserSuite) SetupTest() {
	if s.fns.SetupTest != nil {
		s.fns.SetupTest()
	}
}
