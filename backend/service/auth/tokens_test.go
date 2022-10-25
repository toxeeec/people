package auth

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/token"
)

func (suite *AuthSuite) TestNewTokens() {
	var user people.AuthUser
	gofakeit.Struct(&user)
	id, _ := suite.us.Create(user)

	tokens, err := suite.as.NewTokens(id)
	assert.NotEmpty(suite.T(), tokens)
	assert.NoError(suite.T(), err)
	var rt people.RefreshToken
	suite.db.Get(&rt, "SELECT token_id, value, user_id FROM token WHERE user_id = $1", id)
	assert.Equal(suite.T(), rt.Value, tokens.RefreshToken)
	assert.Equal(suite.T(), rt.UserID, id)
}

func (suite *AuthSuite) TestUpdateRefreshToken() {
	var user people.AuthUser
	gofakeit.Struct(&user)
	id, _ := suite.us.Create(user)
	tokens, _ := suite.as.NewTokens(id)
	rt, _ := token.ParseRefreshToken(tokens.RefreshToken)

	expected, err := suite.as.UpdateRefreshToken(rt.UserID, rt.ID)
	assert.NotEmpty(suite.T(), expected)
	assert.NoError(suite.T(), err)

	var actual people.RefreshToken
	suite.db.Get(&actual, "SELECT token_id, value, user_id FROM token WHERE user_id = $1", id)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *AuthSuite) TestCheckRefreshToken() {
	var user people.AuthUser
	gofakeit.Struct(&user)
	id, _ := suite.us.Create(user)
	tokens, _ := suite.as.NewTokens(id)
	rt1, _ := token.ParseRefreshToken(tokens.RefreshToken)
	rt2, _ := token.NewRefreshToken(1, nil)

	assert.True(suite.T(), suite.as.CheckRefreshToken(rt1))
	assert.False(suite.T(), suite.as.CheckRefreshToken(rt2))
}
