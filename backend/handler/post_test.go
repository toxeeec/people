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
	var valid people.PostBody
	var emptyContent people.PostBody
	var user people.AuthUser
	gofakeit.Struct(&valid)
	gofakeit.Struct(&user)
	emptyContent.Content = "\t\n \n\t"
	userID, _ := suite.us.Create(user)

	at, _ := token.NewAccessToken(userID)

	tests := map[string]struct {
		body     people.PostBody
		expected int
	}{
		"empty content": {emptyContent, http.StatusBadRequest},
		"valid":         {valid, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(at).WithJsonBody(tc.body).Post("/posts").Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var p people.Post
				result.UnmarshalJsonToObject(&p)
				// trim content for assertion
				valid.TrimContent()

				assert.Equal(suite.T(), valid.Content, p.Content)
			}
		})
	}
}

func (suite *HandlerSuite) TestGetPostsPostID() {
	var expected people.PostBody
	var user people.AuthUser
	gofakeit.Struct(&expected)
	gofakeit.Struct(&user)
	userID, _ := suite.us.Create(user)
	p, _ := suite.ps.Create(userID, expected)

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"unknown id": {p.ID + 1, http.StatusNotFound},
		"valid":      {p.ID, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().Get(fmt.Sprintf("/posts/%d", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var p people.Post
				result.UnmarshalJsonToObject(&p)
				// trim content for assertion
				expected.TrimContent()

				assert.Equal(suite.T(), expected.Content, p.Content)
				assert.Equal(suite.T(), user.Handle, p.User.Handle)
			}
		})
	}
}

func (suite *HandlerSuite) TestHandleDeletePost() {
	var expected people.PostBody
	var user people.AuthUser
	gofakeit.Struct(&expected)
	gofakeit.Struct(&user)
	userID, _ := suite.us.Create(user)
	p, _ := suite.ps.Create(userID, expected)

	tests := map[string]struct {
		id       uint
		userID   uint
		expected int
	}{
		"not owned":  {p.ID, userID + 1, http.StatusNotFound},
		"unknown id": {p.ID + 1, userID, http.StatusNotFound},
		"valid":      {p.ID, userID, http.StatusNoContent},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			at, _ := token.NewAccessToken(tc.userID)
			result := testutil.NewRequest().WithJWSAuth(at).Delete(fmt.Sprintf("/posts/%d", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				_, err := suite.ps.Get(p.ID)
				assert.Error(suite.T(), err)
			}
		})
	}
}

func (suite *HandlerSuite) TestGetUsersHandlePosts() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	id1, _ := suite.us.Create(user1)
	suite.us.Create(user2)
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.Create(id1, p)
	}

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"invalid handle": {gofakeit.Username(), 0},
		"0 posts":        {user2.Handle, 0},
		"valid":          {user1.Handle, count},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			pagination := people.NewSeekPagination(nil, nil, nil)
			posts, _ := suite.ps.FromUser(tc.handle, pagination)
			assert.Equal(suite.T(), tc.expected, len(posts.Data))
		})
	}
}
