package user

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *UserSuite) TestFollow() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	var user3 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	gofakeit.Struct(&user3)
	unknownHandle := gofakeit.LetterN(10)
	id1, _ := suite.us.Create(user1)
	id2, _ := suite.us.Create(user2)
	suite.us.Create(user3)
	suite.us.db.MustExec("INSERT INTO follower(user_id, follower_id) VALUES($1, $2)", id2, id1)

	tests := map[string]struct {
		handle string
		valid  bool
	}{
		"unknown handle":   {unknownHandle, false},
		"same user":        {user1.Handle, false},
		"already followed": {user2.Handle, false},
		"valid":            {user3.Handle, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			err := suite.us.Follow(id1, tc.handle)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				var followers uint
				var following uint
				suite.db.Get(&followers, "SELECT followers FROM user_profile WHERE handle = $1", tc.handle)
				suite.db.Get(&following, "SELECT following FROM user_profile WHERE handle = $1", user1.Handle)
				assert.Equal(suite.T(), uint(1), followers)
				assert.Equal(suite.T(), uint(1), following)
			}
		})
	}
}

func (suite *UserSuite) TestUnfollow() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	var user3 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	gofakeit.Struct(&user3)
	unknownHandle := gofakeit.LetterN(10)
	id1, _ := suite.us.Create(user1)
	suite.us.Create(user2)
	suite.us.Create(user3)
	suite.us.Follow(id1, user2.Handle)

	tests := map[string]struct {
		handle string
		valid  bool
	}{
		"unknown handle": {unknownHandle, false},
		"same user":      {user1.Handle, false},
		"not followed":   {user3.Handle, false},
		"valid":          {user2.Handle, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			err := suite.us.Unfollow(id1, tc.handle)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				var followers uint
				var following uint
				suite.db.Get(&followers, "SELECT followers FROM user_profile WHERE handle = $1", tc.handle)
				suite.db.Get(&following, "SELECT following FROM user_profile WHERE handle = $1", user1.Handle)
				assert.Equal(suite.T(), uint(0), followers)
				assert.Equal(suite.T(), uint(0), following)
			}
		})
	}
}

func (suite *UserSuite) TestIsFollowing() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	var user3 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	gofakeit.Struct(&user3)
	unknownHandle := gofakeit.LetterN(10)
	id1, _ := suite.us.Create(user1)
	id2, _ := suite.us.Create(user2)
	suite.us.Create(user3)
	suite.us.db.MustExec("INSERT INTO follower(user_id, follower_id) VALUES($1, $2)", id2, id1)

	tests := map[string]struct {
		handle string
		valid  bool
	}{
		"unknown handle": {unknownHandle, false},
		"not followed":   {user3.Handle, false},
		"valid":          {user2.Handle, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			actual, err := suite.us.IsFollowing(id1, tc.handle)
			assert.Equal(suite.T(), tc.valid, actual)
			assert.NoError(suite.T(), err)
		})
	}
}

func (suite *UserSuite) TestIsFollowed() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	var user3 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	gofakeit.Struct(&user3)
	unknownHandle := gofakeit.LetterN(10)
	id1, _ := suite.us.Create(user1)
	id2, _ := suite.us.Create(user2)
	suite.us.Create(user3)
	suite.us.db.MustExec("INSERT INTO follower(user_id, follower_id) VALUES($1, $2)", id1, id2)

	tests := map[string]struct {
		handle string
		valid  bool
	}{
		"unknown handle": {unknownHandle, false},
		"not followed":   {user3.Handle, false},
		"valid":          {user2.Handle, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			actual, err := suite.us.IsFollowed(id1, tc.handle)
			assert.Equal(suite.T(), tc.valid, actual)
			assert.NoError(suite.T(), err)
		})
	}
}

func (suite *UserSuite) TestFollowing() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	id1, _ := suite.us.Create(user1)
	id2, _ := suite.us.Create(user2)

	var limit uint = 2

	for i := 0; i < 3; i++ {
		var u people.AuthUser
		gofakeit.Struct(&u)
		suite.us.Create(u)
		suite.us.Follow(id1, u.Handle)
	}

	tests := map[string]struct {
		id       uint
		page     uint
		expected int
	}{
		"0 following":  {id2, 1, 0},
		"first page":   {id1, 1, 2},
		"last page":    {id1, 2, 1},
		"page too far": {id1, 3, 0},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			following, _ := suite.us.Following(tc.id, people.NewPagination(&tc.page, &limit))
			assert.Equal(suite.T(), tc.expected, len(following))
		})
	}
}

func (suite *UserSuite) TestFollowers() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	id1, _ := suite.us.Create(user1)
	id2, _ := suite.us.Create(user2)

	var limit uint = 2

	for i := 0; i < 3; i++ {
		var u people.AuthUser
		gofakeit.Struct(&u)
		id, _ := suite.us.Create(u)
		suite.us.Follow(id, user1.Handle)
	}

	tests := map[string]struct {
		id       uint
		page     uint
		expected int
	}{
		"0 followers":  {id2, 1, 0},
		"first page":   {id1, 1, 2},
		"last page":    {id1, 2, 1},
		"page too far": {id1, 3, 0},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			followers, _ := suite.us.Followers(tc.id, people.NewPagination(&tc.page, &limit))
			assert.Equal(suite.T(), tc.expected, len(followers))
		})
	}
}
