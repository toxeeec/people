package handler_test

import (
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *HandlerSuite) TestPostRegister() {
	var valid people.AuthUser
	var takenHandle people.AuthUser
	gofakeit.Struct(&valid)
	gofakeit.Struct(&takenHandle)
	suite.us.Create(takenHandle)

	tests := map[string]struct {
		user     people.AuthUser
		expected int
	}{
		"taken handle": {takenHandle, http.StatusBadRequest},
		"valid":        {valid, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().Post("/register").WithJsonBody(tc.user).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			var tokens people.Tokens
			result.UnmarshalJsonToObject(&tokens)
			if tc.expected < http.StatusBadRequest {
				assert.NotEmpty(suite.T(), tokens)
			} else {
				assert.Empty(suite.T(), tokens)
			}
		})
	}
}

func (suite *HandlerSuite) TestPostLogin() {
	var valid people.AuthUser
	var invalidPassword people.AuthUser
	var unknownHandle people.AuthUser
	gofakeit.Struct(&valid)
	gofakeit.Struct(&invalidPassword)
	gofakeit.Struct(&unknownHandle)
	invalidPassword.Handle = valid.Handle
	suite.us.Create(valid)

	tests := map[string]struct {
		user     people.AuthUser
		expected int
	}{
		"unknown handle":   {unknownHandle, http.StatusUnauthorized},
		"invalid password": {invalidPassword, http.StatusUnauthorized},
		"valid":            {valid, http.StatusOK},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			result := testutil.NewRequest().Post("/login").WithJsonBody(tc.user).Go(suite.T(), suite.e)
			assert.Equal(suite.T(), tc.expected, result.Code())
			var tokens people.Tokens
			result.UnmarshalJsonToObject(&tokens)
			if tc.expected < http.StatusBadRequest {
				assert.NotEmpty(suite.T(), tokens)
			} else {
				assert.Empty(suite.T(), tokens)
			}
		})
	}
}

func (suite *HandlerSuite) newRefreshRequest(rt string, expected int) string {
	time.Sleep(time.Second)
	body := people.Tokens{RefreshToken: rt}
	result := testutil.NewRequest().Post("/refresh").WithJsonBody(body).Go(suite.T(), suite.e)
	assert.Equal(suite.T(), expected, result.Code())
	tokens := people.Tokens{}
	result.UnmarshalJsonToObject(&tokens)

	return tokens.RefreshToken
}

func (suite *HandlerSuite) TestPostRefresh() {
	var user people.AuthUser
	gofakeit.Struct(&user)
	id, _ := suite.us.Create(user)
	tokens, _ := suite.as.NewTokens(id)

	rt := suite.newRefreshRequest(tokens.RefreshToken, http.StatusOK)

	// create new token
	newToken := suite.newRefreshRequest(rt, http.StatusOK)

	// try using the previous token
	suite.newRefreshRequest(rt, http.StatusForbidden)

	// new token is also invalidated now
	suite.newRefreshRequest(newToken, http.StatusForbidden)
}
