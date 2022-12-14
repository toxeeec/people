package repotest

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type PostSuite struct {
	suite.Suite
	repo repository.Post
	ur   repository.User
	fr   repository.Follow
	fns  TestFns
}

func NewPostSuite(pr repository.Post, ur repository.User, fr repository.Follow, fns TestFns) *PostSuite {
	return &PostSuite{repo: pr, ur: ur, fr: fr, fns: fns}
}

func (s *PostSuite) TestCreate() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)

	p, err := s.repo.Create(np, u.ID, nil)
	assert.NoError(s.T(), err)
	actual, _ := s.repo.Get(p.ID)
	assert.Equal(s.T(), p, actual)

	// reply
	reply, err := s.repo.Create(np, u.ID, &p.ID)
	assert.NoError(s.T(), err)
	actual, _ = s.repo.Get(reply.ID)
	assert.Equal(s.T(), reply, actual)
}

func (s *PostSuite) TestGet() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p, _ := s.repo.Create(np, u.ID, nil)

	actual, err := s.repo.Get(p.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), p, actual)
}

func (s *PostSuite) TestDelete() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p, _ := s.repo.Create(np, u.ID, nil)

	s.repo.Delete(p.ID, u.ID)
	_, err := s.repo.Get(p.ID)
	assert.Error(s.T(), err)
}

func (s *PostSuite) TestListUserPosts() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	var posts [3]people.Post
	for i := range posts {
		var np people.NewPost
		gofakeit.Struct(&np)
		posts[i], _ = s.repo.Create(np, u.ID, nil)
	}

	ps, err := s.repo.ListUserPosts(u.ID, pagination.ID{Limit: 10})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), ps, len(posts))

}

func (s *PostSuite) TestListFeed() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	p, _ := s.repo.Create(np, u.ID, nil)

	// replies are excluded
	s.repo.Create(np, u.ID, &p.ID)
	var ids [5]uint
	for i := range ids {
		var au people.AuthUser
		var np people.NewPost
		gofakeit.Struct(&au)
		gofakeit.Struct(&np)
		user, _ := s.ur.Create(au)
		ids[i] = user.ID
		s.fr.Create(ids[i], u.ID)
		s.repo.Create(np, ids[i], nil)
	}

	res, err := s.repo.ListFeed(append(ids[:], u.ID), u.ID, pagination.ID{Limit: 10})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), res, len(ids)+1)
}

func (s *PostSuite) TestListReplies() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	p, _ := s.repo.Create(np, u.ID, nil)
	var users [5]people.User
	for i := range users {
		var au people.AuthUser
		var np people.NewPost
		gofakeit.Struct(&au)
		gofakeit.Struct(&np)
		users[i], _ = s.ur.Create(au)
		s.repo.Create(np, users[i].ID, &p.ID)
	}

	res, err := s.repo.ListReplies(p.ID, pagination.ID{Limit: 10})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), res, len(users))
}

func (s *PostSuite) SetupTest() {
	if s.fns.SetupTest != nil {
		s.fns.SetupTest()
	}
}
