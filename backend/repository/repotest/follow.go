package repotest

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service/user"
)

type FollowSuite struct {
	suite.Suite
	repo repository.Follow
	ur   repository.User
	fns  TestFns
}

func NewFollowSuite(fr repository.Follow, ur repository.User, fns TestFns) *FollowSuite {
	return &FollowSuite{repo: fr, ur: ur, fns: fns}
}

func (s *FollowSuite) TestGetStatusFollowing() {
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.repo.Create(users[0].ID, users[2].ID)

	tests := map[string]struct {
		id        uint
		following bool
	}{
		"not following": {users[1].ID, false},
		"following":     {users[2].ID, true},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			following := s.repo.GetStatusFollowing(tc.id, users[0].ID)
			assert.Equal(s.T(), tc.following, following)
		})
	}
}

func (s *FollowSuite) TestGetStatusFollowed() {
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.repo.Create(users[2].ID, users[0].ID)

	tests := map[string]struct {
		id       uint
		valid    bool
		followed bool
	}{
		"not followed": {users[1].ID, true, false},
		"followed":     {users[2].ID, true, true},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			followed := s.repo.GetStatusFollowed(tc.id, users[0].ID)
			assert.Equal(s.T(), tc.followed, followed)
		})
	}
}

func (s *FollowSuite) TestCreate() {
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	err := s.repo.Create(users[1].ID, users[0].ID)
	assert.NoError(s.T(), err)
	followed := s.repo.GetStatusFollowed(users[1].ID, users[0].ID)
	assert.True(s.T(), followed)

	tests := map[string]struct {
		id  uint
		err error
	}{
		"same user":        {users[0].ID, repository.ErrSameUser},
		"already followed": {users[1].ID, repository.ErrAlreadyFollowed},
		"valid":            {users[2].ID, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			err := s.repo.Create(tc.id, users[0].ID)
			assert.ErrorIs(s.T(), err, tc.err)
			if tc.err == nil {
				assert.True(s.T(), s.repo.GetStatusFollowed(tc.id, users[0].ID))
			}
		})
	}
}

func (s *FollowSuite) TestDelete() {
	var users [2]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.repo.Create(users[1].ID, users[0].ID)

	err := s.repo.Delete(users[1].ID, users[0].ID)
	assert.NoError(s.T(), err)
	assert.False(s.T(), s.repo.GetStatusFollowed(users[1].ID, users[0].ID))
}

func (s *FollowSuite) TestListStatusFollowing() {
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.repo.Create(users[0].ID, users[1].ID)

	fs, err := s.repo.ListStatusFollowing([]uint{users[1].ID, users[2].ID}, users[0].ID)
	assert.NoError(s.T(), err)
	_, u1ok := fs[users[1].ID]
	assert.True(s.T(), u1ok)
	_, u2ok := fs[users[2].ID]
	assert.False(s.T(), u2ok)
}

func (s *FollowSuite) TestListStatusFollowed() {
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.repo.Create(users[1].ID, users[0].ID)

	fs, err := s.repo.ListStatusFollowed([]uint{users[1].ID, users[2].ID}, users[0].ID)
	assert.NoError(s.T(), err)
	_, u1ok := fs[users[1].ID]
	assert.True(s.T(), u1ok)
	_, u2ok := fs[users[2].ID]
	assert.False(s.T(), u2ok)
}

func (s *FollowSuite) TestListFollowing() {
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.repo.Create(users[1].ID, users[0].ID)
	s.repo.Create(users[2].ID, users[0].ID)

	us, err := s.repo.ListFollowing(users[0].ID, &pagination.ID{Limit: 10})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), us, len(users)-1)
}

func (s *FollowSuite) TestListFollowers() {
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
	}
	s.repo.Create(users[0].ID, users[1].ID)
	s.repo.Create(users[0].ID, users[2].ID)

	us, err := s.repo.ListFollowers(users[0].ID, &pagination.ID{Limit: 10})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), us, len(users)-1)
}

func (s *FollowSuite) TestDeleteFollower() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u1, _ := s.ur.Create(au)
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
		s.repo.Create(users[i].ID, u1.ID)
	}

	err := s.repo.DeleteFollower(user.IDs(users[:]))
	assert.NoError(s.T(), err)
	us, _ := s.ur.List(user.IDs(users[:]))
	for _, u := range us {
		assert.Equal(s.T(), uint(0), u.Followers)
	}
}

func (s *FollowSuite) TestDeleteFollowing() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	var users [3]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = s.ur.Create(au)
		s.repo.Create(u.ID, users[i].ID)
	}

	err := s.repo.DeleteFollowing(user.IDs(users[:]))
	assert.NoError(s.T(), err)
	us, _ := s.ur.List(user.IDs(users[:]))
	for _, u := range us {
		assert.Equal(s.T(), uint(0), u.Following)
	}
}

func (s *FollowSuite) SetupTest() {
	if s.fns.SetupTest != nil {
		s.fns.SetupTest()
	}
}
