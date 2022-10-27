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
	var user1 people.AuthUser
	var user2 people.AuthUser
	var user3 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	gofakeit.Struct(&user3)
	id1, _ := suite.us.Create(user1)
	id2, _ := suite.us.Create(user2)
	id3, _ := suite.us.Create(user3)
	suite.us.Follow(id1, user2.Handle)
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.Create(id2, p)
	}

	// not followed by user1
	var p people.PostBody
	gofakeit.Struct(&p)
	suite.ps.Create(id3, p)

	at, _ := token.NewAccessToken(id1)

	result := testutil.NewRequest().WithJWSAuth(at).Get("/me/feed").Go(suite.T(), suite.e)
	assert.Equal(suite.T(), http.StatusOK, result.Code())

	var res people.FeedResponse
	result.UnmarshalJsonToObject(&res)
	assert.Equal(suite.T(), count, len(res.Data))
}
