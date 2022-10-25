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

func (suite *HandlerSuite) TestGetMeFollowersHandle() {
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
	suite.us.Follow(id2, user1.Handle)
	at, _ := token.NewAccessToken(id1)

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"unknown handle": {unknownHandle, http.StatusNotFound},
		"not following":  {user3.Handle, http.StatusNotFound},
		"valid":          {user2.Handle, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(at).Get(fmt.Sprintf("/me/followers/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
		})
	}
}
