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
	p1, _ := suite.ps.Create(userID, post)
	p2, _ := suite.ps.Create(userID, post)
	r, _ := suite.ps.CreateReply(p2.ID, userID, post)

	tests := map[string]struct {
		id     uint
		userID uint
		valid  bool
	}{
		"not owned":     {p1.ID, userID + 1, false},
		"unknown id":    {p1.ID + 3, userID, false},
		"valid (reply)": {r.ID, userID, true},
		"valid":         {p1.ID, userID, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			err := suite.ps.Delete(tc.id, tc.userID)
			assert.Equal(suite.T(), tc.valid, err == nil)
		})
	}
}

func (suite *PostSuite) TestFromUser() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	id1, _ := suite.us.Create(user1)
	suite.us.Create(user2)
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.Create(id1, p)
	}

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"invalid handle": {gofakeit.Username(), 0},
		"0 posts":        {user2.Handle, 0},
		"valid":          {user1.Handle, count},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			pagination := people.NewSeekPagination(nil, nil, nil)
			posts, _ := suite.ps.FromUser(tc.handle, pagination)
			assert.Equal(suite.T(), tc.expected, len(posts.Data))
		})
	}
}

func (suite *PostSuite) TestFeed() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	var user3 people.AuthUser
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	gofakeit.Struct(&user3)
	id1, _ := suite.us.Create(user1)
	id2, _ := suite.us.Create(user2)
	id3, _ := suite.us.Create(user3)
	suite.us.Follow(id1, user2.Handle)
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.Create(id2, p)
	}

	// not followed by user1
	var p people.PostBody
	gofakeit.Struct(&p)
	suite.ps.Create(id3, p)

	pagination := people.NewSeekPagination(nil, nil, nil)
	res, _ := suite.ps.Feed(id1, pagination)
	assert.Equal(suite.T(), count, len(res.Data))
}

func (suite *PostSuite) TestExists() {
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
		"invalid id": {p.ID + 1, false},
		"valid":      {p.ID, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			assert.Equal(suite.T(), tc.valid, suite.ps.Exists(tc.id))
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
