package user

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *UserSuite) TestFollow() {
	suite.us.Follow(suite.id1, suite.user2.Handle)
	u, _ := suite.us.Get(suite.user1.Handle, nil)
	assert.Equal(suite.T(), uint(1), u.Following)
	suite.us.Follow(suite.id4, suite.user1.Handle)

	tests := map[string]struct {
		handle      string
		valid       bool
		isFollowing bool
	}{
		"unknown handle":   {suite.unknownUser.Handle, false, false},
		"same user":        {suite.user1.Handle, false, false},
		"already followed": {suite.user2.Handle, false, false},
		"valid":            {suite.user3.Handle, true, false},
		"is following":     {suite.user4.Handle, true, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			follows, err := suite.us.Follow(suite.id1, tc.handle)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), uint(1), follows.Followers)

				if tc.isFollowing {
					assert.Equal(suite.T(), uint(1), follows.Following)
					assert.True(suite.T(), follows.IsFollowing)
					u, _ := suite.us.Get(suite.user1.Handle, &suite.id4)
					assert.True(suite.T(), u.IsFollowing)
					assert.True(suite.T(), u.IsFollowed)
				} else {
					assert.Equal(suite.T(), uint(0), follows.Following)
					assert.False(suite.T(), follows.IsFollowing)
					u, _ := suite.us.Get(suite.user1.Handle, &suite.id3)
					assert.True(suite.T(), u.IsFollowing)
					assert.False(suite.T(), u.IsFollowed)
				}

			}
		})
	}
}

func (suite *UserSuite) TestUnfollow() {
	suite.us.Follow(suite.id1, suite.user2.Handle)
	u, _ := suite.us.Get(suite.user1.Handle, nil)
	assert.Equal(suite.T(), uint(1), u.Following)

	tests := map[string]struct {
		handle string
		valid  bool
	}{
		"unknown handle": {suite.unknownUser.Handle, false},
		"same user":      {suite.user1.Handle, false},
		"not followed":   {suite.user3.Handle, false},
		"valid":          {suite.user2.Handle, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			follows, err := suite.us.Unfollow(suite.id1, tc.handle)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), uint(0), follows.Followers)
				assert.Equal(suite.T(), uint(0), follows.Following)
				assert.False(suite.T(), follows.IsFollowed)

				u, _ := suite.us.Get(suite.user1.Handle, nil)
				assert.Equal(suite.T(), uint(0), u.Following)
			}
		})
	}
}

func (suite *UserSuite) TestIsFollowing() {
	suite.us.Follow(suite.id1, suite.user2.Handle)

	tests := map[string]struct {
		handle string
		valid  bool
	}{
		"unknown handle": {suite.unknownUser.Handle, false},
		"not followed":   {suite.user3.Handle, false},
		"valid":          {suite.user2.Handle, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			actual, err := suite.us.IsFollowing(suite.id1, tc.handle)
			assert.Equal(suite.T(), tc.valid, actual)
			assert.NoError(suite.T(), err)
		})
	}
}

func (suite *UserSuite) TestIsFollowed() {
	suite.us.Follow(suite.id1, suite.user2.Handle)

	tests := map[string]struct {
		handle string
		valid  bool
	}{
		"unknown handle": {suite.unknownUser.Handle, false},
		"not followed":   {suite.user3.Handle, false},
		"valid":          {suite.user1.Handle, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			actual, err := suite.us.IsFollowed(suite.id2, tc.handle)
			assert.Equal(suite.T(), tc.valid, actual)
			assert.NoError(suite.T(), err)
		})
	}
}

func (suite *UserSuite) TestFollowing() {
	count := 5
	for i := 0; i < 5; i++ {
		var u people.AuthUser
		gofakeit.Struct(&u)
		suite.us.Create(u)
		suite.us.Follow(suite.id1, u.Handle)
	}

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"invalid id":  {suite.id1 + 10, 0},
		"0 following": {suite.id2, 0},
		"valid":       {suite.id1, count},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			following, _ := suite.us.Following(tc.id, new(uint), people.NewPagination[string](nil, nil, nil))
			assert.Equal(suite.T(), tc.expected, len(following.Data))
		})
	}
}

func (suite *UserSuite) TestFollowers() {
	count := 5
	for i := 0; i < count; i++ {
		var u people.AuthUser
		gofakeit.Struct(&u)
		id, _ := suite.us.Create(u)
		suite.us.Follow(id, suite.user1.Handle)
	}

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"invalid id":  {suite.id1 + 10, 0},
		"0 followers": {suite.id2, 0},
		"valid":       {suite.id1, count},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			followers, _ := suite.us.Followers(tc.id, new(uint), people.NewPagination[string](nil, nil, nil))
			assert.Equal(suite.T(), tc.expected, len(followers.Data))
		})
	}
}
