package handler_test

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
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
		"valid":            {suite.user3.Handle, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).Put(fmt.Sprintf("/me/following/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
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

func (suite *HandlerSuite) TestDeleteMeFollowingHandle() {
	suite.us.Follow(suite.id1, suite.user2.Handle)

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"unknown handle": {suite.unknownUser.Handle, http.StatusNotFound},
		"same user":      {suite.user1.Handle, http.StatusNotFound},
		"not followed":   {suite.user3.Handle, http.StatusNotFound},
		"valid":          {suite.user2.Handle, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).Delete(fmt.Sprintf("/me/following/%s", tc.handle)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
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
