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

func (suite *HandlerSuite) TestPostPosts() {
	var emptyContent people.PostBody
	emptyContent.Content = "\t\n \n\t"

	tests := map[string]struct {
		body     people.PostBody
		expected int
	}{
		"empty content": {emptyContent, http.StatusBadRequest},
		"valid":         {suite.postBody1, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).WithJsonBody(tc.body).Post("/posts").Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var p people.Post
				result.UnmarshalJsonToObject(&p)
				// trim content for assertion
				suite.postBody1.TrimContent()

				assert.Equal(suite.T(), suite.postBody1.Content, p.Content)
			}
		})
	}
}

func (suite *HandlerSuite) TestGetPostsPostID() {
	suite.ps.Like(suite.post2.ID, suite.id1)

	tests := map[string]struct {
		id       uint
		expected int
		liked    bool
	}{
		"unknown id": {suite.post1.ID + 5, http.StatusNotFound, false},
		"valid":      {suite.post1.ID, http.StatusOK, false},
		"liked":      {suite.post2.ID, http.StatusOK, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			var result *testutil.CompletedRequest
			if tc.liked {
				result = testutil.NewRequest().Get(fmt.Sprintf("/posts/%d", tc.id)).WithJWSAuth(suite.at1).Go(suite.T(), suite.e)
			} else {
				result = testutil.NewRequest().Get(fmt.Sprintf("/posts/%d", tc.id)).Go(suite.T(), suite.e)
			}
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var p people.Post
				result.UnmarshalJsonToObject(&p)
				if tc.liked {
					suite.postBody2.TrimContent()
					assert.Equal(suite.T(), suite.postBody2.Content, p.Content)
					assert.Equal(suite.T(), suite.user1.Handle, p.User.Handle)
					assert.True(suite.T(), p.IsLiked)
				} else {
					suite.postBody1.TrimContent()
					assert.Equal(suite.T(), suite.postBody1.Content, p.Content)
					assert.Equal(suite.T(), suite.user1.Handle, p.User.Handle)
					assert.False(suite.T(), p.IsLiked)
				}
			}
		})
	}
}

func (suite *HandlerSuite) TestHandleDeletePost() {
	tests := map[string]struct {
		id       uint
		userID   uint
		expected int
	}{
		"not owned":  {suite.post1.ID, suite.id1 + 5, http.StatusNotFound},
		"unknown id": {suite.post1.ID + 5, suite.id1, http.StatusNotFound},
		"valid":      {suite.post1.ID, suite.id1, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			at, _ := token.NewAccessToken(tc.userID)
			result := testutil.NewRequest().WithJWSAuth(at).Delete(fmt.Sprintf("/posts/%d", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				_, err := suite.ps.Get(suite.post1.ID, nil)
				assert.Error(suite.T(), err)
			}
		})
	}
}

func (suite *HandlerSuite) TestGetUsersHandlePosts() {
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.Create(suite.id2, p)
	}

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"invalid handle": {gofakeit.Username(), 0},
		"0 posts":        {suite.user3.Handle, 0},
		"valid":          {suite.user2.Handle, count},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			pagination := people.NewPagination[uint](nil, nil, nil)
			posts, _ := suite.ps.FromUser(tc.handle, nil, pagination)
			assert.Equal(suite.T(), tc.expected, len(posts.Data))
		})
	}
}
