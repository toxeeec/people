package handler_test

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *HandlerSuite) TestPutMeFollowingHandle() {
	suite.us.Follow(suite.id1, suite.user2.Handle)

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"unknown handle":   {suite.unknownUser.Handle, http.StatusNotFound},
		"same user":        {suite.user1.Handle, http.StatusNotFound},
		"already followed": {suite.user2.Handle, http.StatusConflict},
		"valid":            {suite.user3.Handle, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).Put(fmt.Sprintf("/me/following/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var follows people.Follows
				result.UnmarshalJsonToObject(&follows)
				assert.Equal(suite.T(), uint(1), follows.Followers)
				assert.Equal(suite.T(), uint(0), follows.Following)
			}
		})
	}
}

func (suite *HandlerSuite) TestDeleteMeFollowingHandle() {
	suite.us.Follow(suite.id1, suite.user2.Handle)

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"unknown handle": {suite.unknownUser.Handle, http.StatusNotFound},
		"same user":      {suite.user1.Handle, http.StatusNotFound},
		"not followed":   {suite.user3.Handle, http.StatusNotFound},
		"valid":          {suite.user2.Handle, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).Delete(fmt.Sprintf("/me/following/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var follows people.Follows
				result.UnmarshalJsonToObject(&follows)
				assert.Equal(suite.T(), uint(0), follows.Followers)
				assert.Equal(suite.T(), uint(0), follows.Following)
			}
		})
	}
}

func (suite *HandlerSuite) TestGetMeFollowingHandle() {
	suite.us.Follow(suite.id1, suite.user2.Handle)

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"unknown handle": {suite.unknownUser.Handle, http.StatusNotFound},
		"not followed":   {suite.user3.Handle, http.StatusNotFound},
		"valid":          {suite.user2.Handle, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).Get(fmt.Sprintf("/me/following/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
		})
	}
}
