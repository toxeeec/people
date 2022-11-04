package user

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *UserSuite) TestFollow() {
	suite.us.Follow(suite.id1, suite.user2.Handle)

	tests := map[string]struct {
		handle string
		valid  bool
	}{
		"unknown handle":   {suite.unknownUser.Handle, false},
		"same user":        {suite.user1.Handle, false},
		"already followed": {suite.user2.Handle, false},
		"valid":            {suite.user3.Handle, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			err := suite.us.Follow(suite.id1, tc.handle)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				var followers uint
				var following uint
				suite.db.Get(&followers, "SELECT followers FROM user_profile WHERE handle = $1", tc.handle)
				suite.db.Get(&following, "SELECT following FROM user_profile WHERE handle = $1", suite.user1.Handle)
				assert.Equal(suite.T(), uint(1), followers)
				assert.Equal(suite.T(), uint(2), following)
			}
		})
	}
}

func (suite *UserSuite) TestUnfollow() {
	suite.us.Follow(suite.id1, suite.user2.Handle)

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
			err := suite.us.Unfollow(suite.id1, tc.handle)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				var followers uint
				var following uint
				suite.db.Get(&followers, "SELECT followers FROM user_profile WHERE handle = $1", tc.handle)
				suite.db.Get(&following, "SELECT following FROM user_profile WHERE handle = $1", suite.user1.Handle)
				assert.Equal(suite.T(), uint(0), followers)
				assert.Equal(suite.T(), uint(0), following)
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
	var limit uint = 2
	for i := 0; i < 3; i++ {
		var u people.AuthUser
		gofakeit.Struct(&u)
		suite.us.Create(u)
		suite.us.Follow(suite.id1, u.Handle)
	}

	tests := map[string]struct {
		id       uint
		page     uint
		expected int
	}{
		"0 following":  {suite.id2, 1, 0},
		"first page":   {suite.id1, 1, 2},
		"last page":    {suite.id1, 2, 1},
		"page too far": {suite.id1, 3, 0},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			following, _ := suite.us.Following(tc.id, people.NewPagination(&tc.page, &limit))
			assert.Equal(suite.T(), tc.expected, len(following))
		})
	}
}

func (suite *UserSuite) TestFollowers() {
	var limit uint = 2
	for i := 0; i < 3; i++ {
		var u people.AuthUser
		gofakeit.Struct(&u)
		id, _ := suite.us.Create(u)
		suite.us.Follow(id, suite.user1.Handle)
	}

	tests := map[string]struct {
		id       uint
		page     uint
		expected int
	}{
		"0 followers":  {suite.id2, 1, 0},
		"first page":   {suite.id1, 1, 2},
		"last page":    {suite.id1, 2, 1},
		"page too far": {suite.id1, 3, 0},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			followers, _ := suite.us.Followers(tc.id, people.NewPagination(&tc.page, &limit))
			assert.Equal(suite.T(), tc.expected, len(followers))
		})
	}
}
