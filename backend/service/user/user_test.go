package user

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
)

type UserSuite struct {
	suite.Suite
	db *sqlx.DB
	us service
}

func (suite *UserSuite) TestCreate() {
	var user people.AuthUser
	gofakeit.Struct(&user)
	id, _ := suite.us.Create(user)

	rows, _ := suite.db.Queryx("SELECT user_id, handle FROM user_profile")
	for rows.Next() {
		var actual people.User
		rows.StructScan(&actual)
		assert.Equal(suite.T(), id, actual.ID)
		assert.Equal(suite.T(), user.Handle, actual.Handle)
	}
}

func (suite *UserSuite) TestExists() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	suite.us.Create(user1)

	assert.True(suite.T(), suite.us.Exists(user1.Handle))
	assert.False(suite.T(), suite.us.Exists(user2.Handle))
}

func (suite *UserSuite) TestDelete() {
	var user1 people.AuthUser
	gofakeit.Struct(&user1)
	suite.us.Create(user1)

	suite.us.Delete(user1.Handle)
	assert.False(suite.T(), suite.us.Exists(user1.Handle))
}

func (suite *UserSuite) TestGet() {
	var expected people.AuthUser
	gofakeit.Struct(&expected)
	suite.us.Create(expected)

	actual, err := suite.us.Get(expected.Handle)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected.Handle, actual.Handle)

	_, err = suite.us.Get(gofakeit.LetterN(10))
	assert.Error(suite.T(), err)
}

func (suite *UserSuite) SetupSuite() {
	db, err := people.PostgresConnect()
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.db = db
	suite.us = service{db}
}

func (suite *UserSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *UserSuite) SetupTest() {
	suite.db.MustExec("TRUNCATE user_profile CASCADE")
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
