package handler_test

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *HandlerSuite) TestHandleRegister() {
	var valid people.AuthUser
	var takenHandle people.AuthUser
	gofakeit.Struct(&valid)
	gofakeit.Struct(&takenHandle)
	takenHandle.Handle = valid.Handle

	result := testutil.NewRequest().Post("/register").WithJsonBody(valid).Go(suite.T(), suite.e)
	assert.Equal(suite.T(), http.StatusOK, result.Code())
	assert.True(suite.T(), suite.us.Exists(valid.Handle))
	var tokens people.Tokens
	result.UnmarshalJsonToObject(&tokens)
	assert.NotEmpty(suite.T(), tokens)

	result = testutil.NewRequest().Post("/register").WithJsonBody(takenHandle).Go(suite.T(), suite.e)
	assert.Equal(suite.T(), http.StatusBadRequest, result.Code())
	tokens = people.Tokens{}
	result.UnmarshalJsonToObject(&tokens)
	assert.Empty(suite.T(), tokens)
}

func (suite *HandlerSuite) TestHandleLogin() {
	var valid people.AuthUser
	var invalidPassword people.AuthUser
	var unknownHandle people.AuthUser
	gofakeit.Struct(&valid)
	gofakeit.Struct(&invalidPassword)
	gofakeit.Struct(&unknownHandle)
	invalidPassword.Handle = valid.Handle
	suite.us.Create(valid)

	result := testutil.NewRequest().Post("/login").WithJsonBody(valid).Go(suite.T(), suite.e)
	assert.Equal(suite.T(), http.StatusOK, result.Code())
	var tokens people.Tokens
	result.UnmarshalJsonToObject(&tokens)
	assert.NotEmpty(suite.T(), tokens)

	result = testutil.NewRequest().Post("/login").WithJsonBody(invalidPassword).Go(suite.T(), suite.e)
	assert.Equal(suite.T(), http.StatusUnauthorized, result.Code())
	tokens = people.Tokens{}
	result.UnmarshalJsonToObject(&tokens)
	assert.Empty(suite.T(), tokens)

	result = testutil.NewRequest().Post("/login").WithJsonBody(unknownHandle).Go(suite.T(), suite.e)
	assert.Equal(suite.T(), http.StatusUnauthorized, result.Code())
	tokens = people.Tokens{}
	result.UnmarshalJsonToObject(&tokens)
	assert.Empty(suite.T(), tokens)
}
