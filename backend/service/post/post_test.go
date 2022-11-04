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
	ps        service
	us        people.UserService
	user1     people.AuthUser
	user2     people.AuthUser
	user3     people.AuthUser
	user1ID   uint
	user2ID   uint
	user3ID   uint
	post1     people.Post
	post2     people.Post
	post3     people.Post
	postBody1 people.PostBody
	postBody2 people.PostBody
	postBody3 people.PostBody
	replyBody people.PostBody
}

func (suite *PostSuite) TestCreate() {
	rows, err := suite.ps.db.Queryx(`SELECT post_id, content, user_id AS "user.user_id" FROM post WHERE post_id = $1`, suite.post1.ID)
	assert.NoError(suite.T(), err)
	for rows.Next() {
		var actual people.Post
		rows.StructScan(&actual)
		assert.Equal(suite.T(), suite.post1.ID, actual.ID)
		assert.Equal(suite.T(), suite.postBody1.Content, actual.Content)
		assert.Equal(suite.T(), suite.user1ID, *actual.User.ID)
	}
}

func (suite *PostSuite) TestGet() {
	tests := map[string]struct {
		id    uint
		valid bool
	}{
		"unknown id": {suite.post1.ID + 5, false},
		"valid":      {suite.post1.ID, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			p, err := suite.ps.Get(tc.id)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), suite.post1.Content, p.Content)
				assert.Equal(suite.T(), suite.user1.Handle, p.User.Handle)
			}
		})
	}
}

func (suite *PostSuite) TestDelete() {
	r, _ := suite.ps.CreateReply(suite.post2.ID, suite.user1ID, suite.replyBody)

	tests := map[string]struct {
		id     uint
		userID uint
		valid  bool
	}{
		"not owned":     {suite.post1.ID, suite.user1ID + 5, false},
		"unknown id":    {suite.post1.ID + 5, suite.user1ID, false},
		"valid (reply)": {r.ID, suite.user1ID, true},
		"valid":         {suite.post1.ID, suite.user1ID, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			err := suite.ps.Delete(tc.id, tc.userID)
			assert.Equal(suite.T(), tc.valid, err == nil)
		})
	}
}

func (suite *PostSuite) TestFromUser() {
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.Create(suite.user2ID, p)
	}

	tests := map[string]struct {
		handle   string
		expected int
	}{
		"invalid handle": {gofakeit.Username(), 0},
		"0 posts":        {suite.user3.Handle, 0},
		"valid":          {suite.user2.Handle, count},
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
	suite.us.Follow(suite.user1ID, suite.user2.Handle)
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.Create(suite.user2ID, p)
	}

	// not followed by user1
	var p people.PostBody
	gofakeit.Struct(&p)
	suite.ps.Create(suite.user3ID, p)

	pagination := people.NewSeekPagination(nil, nil, nil)
	res, _ := suite.ps.Feed(suite.user1ID, pagination)
	assert.Equal(suite.T(), count, len(res.Data))
}

func (suite *PostSuite) TestExists() {
	tests := map[string]struct {
		id    uint
		valid bool
	}{
		"invalid id": {suite.post1.ID + 5, false},
		"valid":      {suite.post1.ID, true},
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
	gofakeit.Struct(&suite.user1)
	gofakeit.Struct(&suite.user2)
	gofakeit.Struct(&suite.user3)
	suite.user1ID, _ = suite.us.Create(suite.user1)
	suite.user2ID, _ = suite.us.Create(suite.user2)
	suite.user3ID, _ = suite.us.Create(suite.user3)
}

func (suite *PostSuite) TearDownSuite() {
	suite.ps.db.Close()
}

func (suite *PostSuite) SetupTest() {
	suite.ps.db.MustExec("TRUNCATE post CASCADE")
	gofakeit.Struct(&suite.postBody1)
	gofakeit.Struct(&suite.postBody2)
	gofakeit.Struct(&suite.replyBody)
	suite.post1, _ = suite.ps.Create(suite.user1ID, suite.postBody1)
	suite.post2, _ = suite.ps.Create(suite.user1ID, suite.postBody2)
}

func TestPostSuite(t *testing.T) {
	suite.Run(t, new(PostSuite))
}
