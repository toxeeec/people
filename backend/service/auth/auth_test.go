package auth

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/user"
)

type AuthSuite struct {
	suite.Suite
	db *sqlx.DB
	us people.UserService
	as service
}

func (suite *AuthSuite) TestVerifyCredentials() {
	var valid people.AuthUser
	var invalidPassword people.AuthUser
	var unknownHandle people.AuthUser
	gofakeit.Struct(&valid)
	gofakeit.Struct(&invalidPassword)
	gofakeit.Struct(&unknownHandle)
	invalidPassword.Handle = valid.Handle
	id1, _ := suite.us.Create(valid)

	id, err := suite.as.VerifyCredentials(valid)
	assert.Equal(suite.T(), id1, id)
	assert.NoError(suite.T(), err)

	_, err = suite.as.VerifyCredentials(invalidPassword)
	assert.Error(suite.T(), err)

	_, err = suite.as.VerifyCredentials(unknownHandle)
	assert.Error(suite.T(), err)
}

func (suite *AuthSuite) SetupSuite() {
	db, err := people.PostgresConnect()
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.db = db
	us := user.NewService(db)
	suite.us = us
	suite.as = service{db, us}
}

func (suite *AuthSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *AuthSuite) SetupTest() {
	suite.db.MustExec("TRUNCATE user_profile CASCADE")
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}
