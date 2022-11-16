package handler_test

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *HandlerSuite) TestGetUsersHandle() {
	suite.us.Follow(suite.id2, suite.user1.Handle)

	tests := map[string]struct {
		handle      string
		valid       bool
		isFollowing bool
		expected    int
	}{
		"unknown handle": {suite.unknownUser.Handle, false, false, http.StatusNotFound},
		"is following":   {suite.user2.Handle, true, true, http.StatusOK},
		"valid":          {suite.user3.Handle, true, false, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).Get(fmt.Sprintf("/users/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var u people.User
				result.UnmarshalJsonToObject(&u)
				assert.Equal(suite.T(), tc.handle, u.Handle)
				assert.Equal(suite.T(), tc.isFollowing, u.IsFollowing)
			}
		})
	}
}
