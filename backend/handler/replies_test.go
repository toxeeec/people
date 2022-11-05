package handler_test

import (
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *HandlerSuite) TestPostPostsPostIDReplies() {
	var emptyContent people.PostBody
	emptyContent.Content = "\t\n \n\t"

	tests := map[string]struct {
		id       uint
		body     people.PostBody
		expected int
	}{
		"invalid id":    {suite.post1.ID + 5, suite.postBody1, http.StatusNotFound},
		"empty content": {suite.post1.ID, emptyContent, http.StatusBadRequest},
		"valid":         {suite.post1.ID, suite.postBody1, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(suite.at1).WithJsonBody(tc.body).Post(fmt.Sprintf("/posts/%d/replies", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var actual people.Post
				result.UnmarshalJsonToObject(&actual)
				// trim content for assertion
				suite.postBody1.TrimContent()

				assert.Equal(suite.T(), suite.postBody1.Content, actual.Content)
				assert.Equal(suite.T(), suite.post1.ID, uint(actual.RepliesTo.Int32))
			}
		})
	}
}

func (suite *HandlerSuite) TestGetPostsPostIDReplies() {
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.CreateReply(suite.post1.ID, suite.id1, p)
	}

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"invalid id": {suite.post1.ID + 5, 0},
		"0 replies":  {suite.post2.ID, 0},
		"valid":      {suite.post1.ID, count},
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
