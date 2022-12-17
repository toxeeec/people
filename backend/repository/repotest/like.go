package repotest

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type LikeSuite struct {
	suite.Suite
	repo repository.Like
	pr   repository.Post
	ur   repository.User
	fns  TestFns
}

func NewLikeSuite(lr repository.Like, pr repository.Post, ur repository.User, fns TestFns) *LikeSuite {
	return &LikeSuite{repo: lr, pr: pr, ur: ur, fns: fns}
}

func (s *LikeSuite) TestStatus() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p1, _ := s.pr.Create(np, u.ID, nil)
	p2, _ := s.pr.Create(np, u.ID, nil)
	s.repo.Create(p2.ID, u.ID)

	tests := map[string]struct {
		postID uint
		userID uint
		liked  bool
	}{
		"not liked": {p1.ID, u.ID, false},
		"liked":     {p2.ID, u.ID, true},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			ls := s.repo.Status(tc.postID, tc.userID)
			assert.Equal(s.T(), tc.liked, ls.IsLiked)
		})
	}
}

func (s *LikeSuite) TestCreate() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p1, _ := s.pr.Create(np, u.ID, nil)
	p2, _ := s.pr.Create(np, u.ID, nil)
	err := s.repo.Create(p1.ID, u.ID)
	assert.NoError(s.T(), err)
	ls := s.repo.Status(p1.ID, u.ID)
	assert.True(s.T(), ls.IsLiked)

	tests := map[string]struct {
		postID uint
		userID uint
		err    *error
	}{
		"invalid post id": {p2.ID + 5, u.ID, &repository.ErrPostNotFound},
		"already liked":   {p1.ID, u.ID, &repository.ErrAlreadyLiked},
		"valid":           {p2.ID, u.ID, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			err := s.repo.Create(tc.postID, tc.userID)
			if tc.err != nil {
				assert.ErrorIs(s.T(), err, *tc.err)
			} else {
				assert.NoError(s.T(), err)
				ls := s.repo.Status(tc.postID, tc.userID)
				assert.True(s.T(), ls.IsLiked)
			}
		})
	}
}

func (s *LikeSuite) TestDelete() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p1, _ := s.pr.Create(np, u.ID, nil)
	p2, _ := s.pr.Create(np, u.ID, nil)
	s.repo.Create(p2.ID, u.ID)

	tests := map[string]struct {
		postID uint
		userID uint
		valid  bool
	}{
		"invalid post id": {p2.ID + 5, u.ID, false},
		"not liked":       {p1.ID, u.ID, false},
		"valid":           {p2.ID, u.ID, true},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			err := s.repo.Delete(tc.postID, tc.userID)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.NoError(s.T(), err)
				ls := s.repo.Status(tc.postID, tc.userID)
				assert.False(s.T(), ls.IsLiked)
			}
		})
	}
}

func (s *LikeSuite) TestListUsers() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p, _ := s.pr.Create(np, u.ID, nil)
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
		s.repo.Create(p.ID, users[i].ID)
	}

	us, _ := s.repo.ListPostLikes(p.ID, pagination.ID{Limit: 10})
	assert.Len(s.T(), us.Data, len(users))
}

func (s *LikeSuite) TestListStatusLiked() {
	var au people.AuthUser
	u, _ := s.ur.Create(au)
	var posts [3]people.Post
	for i := range posts {
		var np people.NewPost
		gofakeit.Struct(&np)
		posts[i], _ = s.pr.Create(np, u.ID, nil)
	}
	s.repo.Create(posts[0].ID, u.ID)

	fs, err := s.repo.ListStatusLiked([]uint{posts[0].ID, posts[1].ID, posts[2].ID}, u.ID)
	assert.NoError(s.T(), err)
	_, u1ok := fs[posts[0].ID]
	assert.True(s.T(), u1ok)
	_, u2ok := fs[posts[1].ID]
	assert.False(s.T(), u2ok)
	_, u3ok := fs[posts[2].ID]
	assert.False(s.T(), u3ok)
}

func (s *LikeSuite) TestListUserLikes() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	var posts [3]people.Post
	for i := range posts {
		var np people.NewPost
		gofakeit.Struct(&np)
		posts[i], _ = s.pr.Create(np, u.ID, nil)
		s.repo.Create(posts[i].ID, u.ID)
	}

	ps, err := s.repo.ListUserLikes(u.ID, pagination.ID{Limit: 10})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), ps, len(posts))
}

func (s *LikeSuite) SetupTest() {
	if s.fns.SetupTest != nil {
		s.fns.SetupTest()
	}
}
