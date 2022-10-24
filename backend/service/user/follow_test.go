package user

import (
	"fmt"

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
	suite.us.db.MustExec(fmt.Sprintf("INSERT INTO follower(user_id, follower_id) VALUES(%d, %d)", id2, id1))

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
				user1, _ := suite.us.Get(user1.Handle)
				followed, _ := suite.us.Get(tc.handle)
				assert.Equal(suite.T(), uint(1), followed.Followers)
				assert.Equal(suite.T(), uint(1), user1.Following)
			}
		})
	}
}