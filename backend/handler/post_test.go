package handler_test

import (
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
