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

func (suite *HandlerSuite) TestPostPostsPostIDReplies() {
	var valid people.PostBody
	var emptyContent people.PostBody
	var user people.AuthUser
	gofakeit.Struct(&valid)
	gofakeit.Struct(&user)
	emptyContent.Content = "\t\n \n\t"
	userID, _ := suite.us.Create(user)
	p, _ := suite.ps.Create(userID, valid)

	at, _ := token.NewAccessToken(userID)

	tests := map[string]struct {
		id       uint
		body     people.PostBody
		expected int
	}{
		"invalid id":    {p.ID + 2, valid, http.StatusNotFound},
		"empty content": {p.ID, emptyContent, http.StatusBadRequest},
		"valid":         {p.ID, valid, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(at).WithJsonBody(tc.body).Post(fmt.Sprintf("/posts/%d/replies", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var actual people.Post
				result.UnmarshalJsonToObject(&actual)
				// trim content for assertion
				valid.TrimContent()

				assert.Equal(suite.T(), valid.Content, actual.Content)
				assert.Equal(suite.T(), p.ID, uint(actual.RepliesTo.Int32))
			}
		})
	}
}

func (suite *HandlerSuite) TestGetPostsPostIDReplies() {
	var user people.AuthUser
	var post1 people.PostBody
	var post2 people.PostBody
	gofakeit.Struct(&user)
	gofakeit.Struct(&post1)
	gofakeit.Struct(&post2)
	userID, _ := suite.us.Create(user)
	p1, _ := suite.ps.Create(userID, post1)
	p2, _ := suite.ps.Create(userID, post2)

	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.CreateReply(p1.ID, userID, p)
	}

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"invalid id": {p1.ID + 2, 0},
		"0 replies":  {p2.ID, 0},
		"valid":      {p1.ID, count},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().Get(fmt.Sprintf("/posts/%d/replies", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), http.StatusOK, result.Code())
			var posts people.Posts
			result.UnmarshalJsonToObject(&posts)
			assert.Equal(suite.T(), tc.expected, len(posts.Data))
		})
	}
}
