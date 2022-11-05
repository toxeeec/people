package handler_test

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/token"
)

func (suite *HandlerSuite) TestGetMeFeed() {
	suite.us.Follow(suite.id1, suite.user2.Handle)
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.Create(suite.id2, p)
	}

	// not followed by user1
	var p people.PostBody
	gofakeit.Struct(&p)
	suite.ps.Create(suite.id3, p)

	at, _ := token.NewAccessToken(suite.id1)

	result := testutil.NewRequest().WithJWSAuth(at).Get("/me/feed").Go(suite.T(), suite.e)
	assert.Equal(suite.T(), http.StatusOK, result.Code())

	var res people.PaginationResult[people.Post, uint]
	result.UnmarshalJsonToObject(&res)
	assert.Equal(suite.T(), count, len(res.Data))
}
