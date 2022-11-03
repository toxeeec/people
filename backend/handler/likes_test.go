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

func (suite *HandlerSuite) TestPutPostsPostIDLikes() {
	var post people.PostBody
	var user people.AuthUser
	gofakeit.Struct(&post)
	gofakeit.Struct(&user)
	userID, _ := suite.us.Create(user)
	p1, _ := suite.ps.Create(userID, post)
	p2, _ := suite.ps.Create(userID, post)

	suite.ps.Like(p2.ID, userID)

	at, _ := token.NewAccessToken(userID)

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"invalid id":    {p1.ID + 2, http.StatusNotFound},
		"already liked": {p2.ID, http.StatusConflict},
		"valid":         {p1.ID, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().WithJWSAuth(at).Put(fmt.Sprintf("/posts/%d/likes", tc.id)).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			if tc.expected < http.StatusBadRequest {
				var l people.Likes
				result.UnmarshalJsonToObject(&l)
				assert.Equal(suite.T(), uint(1), l.Likes)
			}
		})
	}
}
