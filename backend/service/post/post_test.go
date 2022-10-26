package post

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/user"
)

type PostSuite struct {
	suite.Suite
	ps service
	us people.UserService
}

func (suite *PostSuite) TestCreate() {
	var post people.PostBody
	var user people.AuthUser
	gofakeit.Struct(&post)
	gofakeit.Struct(&user)
	userID, _ := suite.us.Create(user)
	p, _ := suite.ps.Create(userID, post)

	rows, err := suite.ps.db.Queryx(`SELECT post_id, content, user_id AS "user.user_id" FROM post`)
	assert.NoError(suite.T(), err)
	for rows.Next() {
		var actual people.Post
		rows.StructScan(&actual)
		assert.Equal(suite.T(), p.ID, actual.ID)
		assert.Equal(suite.T(), post.Content, actual.Content)
		assert.Equal(suite.T(), userID, *actual.User.ID)
	}
}

func (suite *PostSuite) SetupSuite() {
	db, err := people.PostgresConnect()
	if err != nil {
		suite.T().Fatal(err.Error())
	}

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	suite.us = user.NewService(db)
	suite.ps = service{db}
}

func (suite *PostSuite) TearDownSuite() {
	suite.ps.db.Close()
}

func (suite *PostSuite) SetupTest() {
	suite.ps.db.MustExec("TRUNCATE post CASCADE")
	suite.ps.db.MustExec("TRUNCATE user_profile CASCADE")
}

func TestPostSuite(t *testing.T) {
	suite.Run(t, new(PostSuite))
}
