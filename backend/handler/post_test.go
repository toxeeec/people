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
	tests := map[string]struct {
		id       uint
		expected int
	}{
		"unknown id": {suite.post1.ID + 5, http.StatusNotFound},
		"valid":      {suite.post1.ID, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().Get(fmt.Sprintf("/posts/%d", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var p people.Post
				result.UnmarshalJsonToObject(&p)
				// trim content for assertion
				suite.postBody1.TrimContent()

				assert.Equal(suite.T(), suite.postBody1.Content, p.Content)
				assert.Equal(suite.T(), suite.user1.Handle, p.User.Handle)
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
			posts, err := suite.ps.FromUser(tc.handle, nil, pagination)
			if err != nil {
				println(err.Error())
			}
			assert.Equal(suite.T(), tc.expected, len(posts.Data))
		})
	}
}
