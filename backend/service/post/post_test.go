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

func (suite *PostSuite) TestGet() {
	var post people.PostBody
	var user people.AuthUser
	gofakeit.Struct(&post)
	gofakeit.Struct(&user)
	userID, _ := suite.us.Create(user)
	p, _ := suite.ps.Create(userID, post)

	tests := map[string]struct {
		id    uint
		valid bool
	}{
		"unknown id": {p.ID + 1, false},
		"valid":      {p.ID, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			p, err := suite.ps.Get(tc.id)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), post.Content, p.Content)
				assert.Equal(suite.T(), user.Handle, p.User.Handle)
			}
		})
	}
}

func (suite *PostSuite) TestDelete() {
	var post people.PostBody
	var user people.AuthUser
	gofakeit.Struct(&post)
	gofakeit.Struct(&user)
	userID, _ := suite.us.Create(user)
	p, _ := suite.ps.Create(userID, post)

	tests := map[string]struct {
		id     uint
		userID uint
		valid  bool
	}{
		"not owned":  {p.ID, userID + 1, false},
		"unknown id": {p.ID + 1, userID, false},
		"valid":      {p.ID, userID, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			err := suite.ps.Delete(tc.id, tc.userID)
			assert.Equal(suite.T(), tc.valid, err == nil)
		})
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
