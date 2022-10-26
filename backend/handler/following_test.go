package handler_test

import (
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/token"
)

func (suite *HandlerSuite) TestPutMeFollowingHandle() {
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
	suite.db.MustExec(fmt.Sprintf("INSERT INTO follower(user_id, follower_id) VALUES(%d, %d)", id2, id1))
	at, _ := token.NewAccessToken(id1)

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"unknown handle":   {unknownHandle, http.StatusNotFound},
		"same user":        {user1.Handle, http.StatusNotFound},
		"already followed": {user2.Handle, http.StatusConflict},
		"valid":            {user3.Handle, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(at).Put(fmt.Sprintf("/me/following/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
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

func (suite *HandlerSuite) TestDeleteMeFollowingHandle() {
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
	at, _ := token.NewAccessToken(id1)

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"unknown handle": {unknownHandle, http.StatusConflict},
		"same user":      {user1.Handle, http.StatusConflict},
		"not followed":   {user3.Handle, http.StatusConflict},
		"valid":          {user2.Handle, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(at).Delete(fmt.Sprintf("/me/following/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
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

func (suite *HandlerSuite) TestGetMeFollowingHandle() {
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
	at, _ := token.NewAccessToken(id1)

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"unknown handle": {unknownHandle, http.StatusNotFound},
		"not followed":   {user3.Handle, http.StatusNotFound},
		"valid":          {user2.Handle, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(at).Get(fmt.Sprintf("/me/following/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
		})
	}
}
