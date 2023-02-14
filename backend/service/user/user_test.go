package user_test

import (
	"context"
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
	"github.com/toxeeec/people/backend/service/user"
)

type UserSuite struct {
	suite.Suite
	us user.Service
	ur repository.User
	fr repository.Follow
	lr repository.Like
	pr repository.Post
}

func (s *UserSuite) TestGetUser() {
	var unknownUser people.AuthUser
	gofakeit.Struct(&unknownUser)
	var users [4]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.fr.Create(users[0].ID, users[1].ID)
	s.fr.Create(users[2].ID, users[0].ID)
	s.fr.Create(users[3].ID, users[0].ID)
	s.fr.Create(users[0].ID, users[3].ID)

	tests := map[string]struct {
		handle    string
		id        uint
		valid     bool
		followed  bool
		following bool
	}{
		"unknown handle":    {unknownUser.Handle, 0, false, false, false},
		"not authenticated": {users[0].Handle, 0, true, false, false},
		"following":         {users[1].Handle, users[0].ID, true, false, true},
		"followed":          {users[2].Handle, users[0].ID, true, true, false},
		"both":              {users[3].Handle, users[0].ID, true, true, true},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			u, err := s.us.GetUser(context.Background(), tc.handle, tc.id, tc.id != 0)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(s.T(), tc.handle, u.Handle)
				if tc.id != 0 {
					assert.Equal(s.T(), tc.following, u.Status.IsFollowing)
					assert.Equal(s.T(), tc.followed, u.Status.IsFollowed)
				}
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), people.NotFoundError, *e.Kind)
			}
		})
	}
}

func (s *UserSuite) TestFollow() {
	var unknownUser people.AuthUser
	gofakeit.Struct(&unknownUser)
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.fr.Create(users[1].ID, users[0].ID)

	notFoundError := people.NotFoundError
	conflictError := people.ConflictError

	tests := map[string]struct {
		handle string
		valid  bool
		kind   *people.ErrorKind
	}{
		"unknown handle":   {unknownUser.Handle, false, &notFoundError},
		"same user":        {users[0].Handle, false, &notFoundError},
		"already followed": {users[1].Handle, false, &conflictError},
		"valid":            {users[2].Handle, true, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			u, err := s.us.Follow(context.Background(), tc.handle, users[0].ID)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(s.T(), tc.handle, u.Handle)
				assert.True(s.T(), u.Status.IsFollowed)
				assert.Equal(s.T(), uint(1), u.Followers)
				assert.Equal(s.T(), uint(0), u.Following)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}
		})
	}
}

func (s *UserSuite) TestUnfollow() {
	var unknownUser people.AuthUser
	gofakeit.Struct(&unknownUser)
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.fr.Create(users[2].ID, users[0].ID)

	notFoundError := people.NotFoundError

	tests := map[string]struct {
		handle string
		valid  bool
		kind   *people.ErrorKind
	}{
		"unknown handle": {unknownUser.Handle, false, &notFoundError},
		"same user":      {users[0].Handle, false, &notFoundError},
		"not followed":   {users[1].Handle, false, &notFoundError},
		"valid":          {users[2].Handle, true, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			u, err := s.us.Unfollow(context.Background(), tc.handle, users[0].ID)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(s.T(), tc.handle, u.Handle)
				assert.False(s.T(), u.Status.IsFollowed)
				assert.Equal(s.T(), uint(0), u.Followers)
				assert.Equal(s.T(), uint(0), u.Following)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}
		})
	}
}

func (s *UserSuite) TestListFollowing() {
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.fr.Create(users[1].ID, users[0].ID)
	s.fr.Create(users[2].ID, users[0].ID)

	tests := map[string]struct {
		handle string
		id     uint
	}{
		"not authenticated": {users[0].Handle, 0},
		"authenticated":     {users[0].Handle, users[0].ID},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			us, err := s.us.ListFollowing(context.Background(), tc.handle, tc.id, tc.id != 0, pagination.HandleParams{})
			assert.NoError(s.T(), err)
			assert.Len(s.T(), us.Data, len(users)-1)
			if tc.id != 0 {
				for _, v := range us.Data {
					assert.True(s.T(), v.Status.IsFollowed)
				}
			}
		})
	}
}

func (s *UserSuite) TestListFollowers() {
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.fr.Create(users[0].ID, users[1].ID)
	s.fr.Create(users[0].ID, users[2].ID)

	tests := map[string]struct {
		handle string
		id     uint
	}{
		"not authenticated": {users[0].Handle, 0},
		"authenticated":     {users[0].Handle, users[0].ID},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			us, err := s.us.ListFollowers(context.Background(), tc.handle, tc.id, tc.id != 0, pagination.HandleParams{})
			assert.NoError(s.T(), err)
			assert.Len(s.T(), us.Data, len(users)-1)
			if tc.id != 0 {
				for _, v := range us.Data {
					assert.True(s.T(), v.Status.IsFollowing)
				}
			}
		})
	}
}

func (s *UserSuite) TestListPostLikes() {
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
		s.lr.Create(p.ID, users[i].ID)
	}

	us, _ := s.us.ListPostLikes(context.Background(), p.ID, 0, false, pagination.HandleParams{})
	assert.Len(s.T(), us.Data, len(users))
}

func (s *UserSuite) TestUpdate() {
	var au1 people.AuthUser
	var au2 people.AuthUser
	gofakeit.Struct(&au1)
	gofakeit.Struct(&au2)
	ar1, _ := s.ur.Create(au1)
	s.ur.Create(au2)

	validationError := people.ValidationError

	tests := map[string]struct {
		handle string
		valid  bool
		kind   *people.ErrorKind
	}{

		"taken handle": {au2.Handle, false, &validationError},
		"valid":        {gofakeit.LetterN(10), true, nil},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			u, err := s.us.Update(ar1.ID, tc.handle)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(s.T(), tc.handle, u.Handle)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)

			}
		})
	}
}

func (s *UserSuite) SetupTest() {
	um := map[uint]people.User{}
	fm := map[inmem.FollowKey]time.Time{}
	lm := map[inmem.LikeKey]struct{}{}
	pm := map[uint]people.Post{}
	v := validator.New()
	s.ur = inmem.NewUserRepository(um)
	s.fr = inmem.NewFollowRepository(fm, um)
	s.lr = inmem.NewLikeRepository(lm, pm, um)
	s.pr = inmem.NewPostRepository(pm)
	s.us = user.NewService(v, s.ur, s.fr, s.lr)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
