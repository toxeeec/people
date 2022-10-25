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

	tests := map[string]struct {
		user  people.AuthUser
		valid bool
	}{
		"unknown handle":   {unknownHandle, false},
		"invalid password": {invalidPassword, false},
		"valid":            {valid, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			id, err := suite.as.VerifyCredentials(tc.user)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), id1, id)
			}
		})
	}
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
