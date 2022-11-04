package auth

import (
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/token"
)

func (suite *AuthSuite) TestNewTokens() {
	tokens, err := suite.as.NewTokens(suite.user1ID)
	assert.NotEmpty(suite.T(), tokens)
	assert.NoError(suite.T(), err)

	var rt people.RefreshToken
	suite.db.Get(&rt, "SELECT token_id, value, user_id FROM token WHERE user_id = $1", suite.user1ID)
	assert.Equal(suite.T(), rt.Value, tokens.RefreshToken)
	assert.Equal(suite.T(), rt.UserID, suite.user1ID)
}

func (suite *AuthSuite) TestUpdateRefreshToken() {
	tokens, _ := suite.as.NewTokens(suite.user1ID)
	rt, _ := token.ParseRefreshToken(tokens.RefreshToken)

	expected, err := suite.as.UpdateRefreshToken(rt.UserID, rt.ID)
	assert.NotEmpty(suite.T(), expected)
	assert.NoError(suite.T(), err)

	var actual people.RefreshToken
	suite.db.Get(&actual, "SELECT token_id, value, user_id FROM token WHERE user_id = $1", suite.user1ID)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *AuthSuite) TestCheckRefreshToken() {
	tokens, _ := suite.as.NewTokens(suite.user1ID)
	rt1, _ := token.ParseRefreshToken(tokens.RefreshToken)
	rt2, _ := token.NewRefreshToken(1, nil)

	assert.True(suite.T(), suite.as.CheckRefreshToken(rt1))
	assert.False(suite.T(), suite.as.CheckRefreshToken(rt2))
}
