package repotest

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type ImageSuite struct {
	suite.Suite
	repo repository.Image
	ur   repository.User
	fns  TestFns
}

func NewImageSuite(ir repository.Image, ur repository.User, fns TestFns) *ImageSuite {
	return &ImageSuite{repo: ir, ur: ur, fns: fns}
}

func (s *ImageSuite) TestCreate() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)

	path := gofakeit.LetterN(40)
	ir, err := s.repo.Create(path, u.ID)
	assert.NoError(s.T(), err)
	assert.NotZero(s.T(), ir.ID)
}

func (s *ImageSuite) TestGet() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	path := gofakeit.LetterN(40)
	ir, _ := s.repo.Create(path, u.ID)

	i, err := s.repo.Get(ir.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), ir.ID, i.ID)
	assert.Equal(s.T(), path, i.Path)
	assert.Equal(s.T(), u.ID, i.UserID)
}

func (s *ImageSuite) TestListUnusedBefore() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	var irs [3]people.ImageResponse
	for i := range irs {
		irs[i], _ = s.repo.Create(gofakeit.LetterN(40), u.ID)
	}
	time.Sleep(1 * time.Second)

	is, err := s.repo.ListUnusedBefore(time.Now())
	assert.NoError(s.T(), err)
	for _, i := range is {
		assert.NotZero(s.T(), i.ID)
	}
}

func (s *ImageSuite) TestDeleteMany() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	var ids [3]uint
	for i := range ids {
		ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}

	s.repo.DeleteMany(ids[:])
	for _, i := range ids {
		_, err := s.repo.Get(i)
		assert.Error(s.T(), err)
	}
}

func (s *ImageSuite) SetupTest() {
	if s.fns.SetupTest != nil {
		s.fns.SetupTest()
	}
}
