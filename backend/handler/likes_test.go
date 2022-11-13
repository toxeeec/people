package handler_test

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *HandlerSuite) TestPutPostsPostIDLikes() {
	suite.ps.Like(suite.post2.ID, suite.id1)

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"invalid id":    {suite.post1.ID + 5, http.StatusNotFound},
		"already liked": {suite.post2.ID, http.StatusConflict},
		"valid":         {suite.post1.ID, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).Put(fmt.Sprintf("/posts/%d/likes", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var l people.Likes
				result.UnmarshalJsonToObject(&l)
				assert.Equal(suite.T(), uint(1), l.Likes)
				assert.True(suite.T(), l.IsLiked)

				p, _ := suite.ps.Get(tc.id, &suite.id1)
				assert.True(suite.T(), p.IsLiked)
			}
		})
	}
}

func (suite *HandlerSuite) TestDeletePostsPostIDLikes() {
	suite.ps.Like(suite.post1.ID, suite.id1)

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"invalid id": {suite.post1.ID + 5, http.StatusNotFound},
		"not liked":  {suite.post2.ID, http.StatusNotFound},
		"valid":      {suite.post1.ID, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).Delete(fmt.Sprintf("/posts/%d/likes", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var l people.Likes
				result.UnmarshalJsonToObject(&l)
				assert.Equal(suite.T(), uint(0), l.Likes)
				assert.False(suite.T(), l.IsLiked)

				p, _ := suite.ps.Get(tc.id, &suite.id1)
				assert.False(suite.T(), p.IsLiked)
			}
		})
	}
}
