package handler_test

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/toxeeec/people/backend/token"
)

func (suite *HandlerSuite) TestGetMeFollowersHandle() {
	suite.us.Follow(suite.id2, suite.user1.Handle)
	at, _ := token.NewAccessToken(suite.id1)

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"unknown handle": {suite.unknownUser.Handle, http.StatusNotFound},
		"not following":  {suite.user3.Handle, http.StatusNotFound},
		"valid":          {suite.user2.Handle, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(at).Get(fmt.Sprintf("/me/followers/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
		})
	}
}
