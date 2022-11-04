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
	db      *sqlx.DB
	us      people.UserService
	as      service
	user1   people.AuthUser
	user1ID uint
}

func (suite *AuthSuite) TestVerifyCredentials() {
	var invalidPassword people.AuthUser
	var unknownHandle people.AuthUser
	gofakeit.Struct(&invalidPassword)
	gofakeit.Struct(&unknownHandle)
	invalidPassword.Handle = suite.user1.Handle

	tests := map[string]struct {
		user  people.AuthUser
		valid bool
	}{
		"unknown handle":   {unknownHandle, false},
		"invalid password": {invalidPassword, false},
		"valid":            {suite.user1, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			id, err := suite.as.VerifyCredentials(tc.user)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), suite.user1ID, id)
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
	gofakeit.Struct(&suite.user1)
	suite.user1ID, _ = suite.us.Create(suite.user1)
}

func (suite *AuthSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *AuthSuite) SetupTest() {
	suite.db.MustExec("TRUNCATE token CASCADE")
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}
