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
			page := uint(1)
			limit := uint(20)
			pagination := people.NewPagination(&page, &limit)
			posts, _ := suite.ps.FromUser(tc.handle, pagination)
			assert.Equal(suite.T(), tc.expected, len(posts))
		})
	}
}

func (suite *PostSuite) TestFeed() {
	var user people.AuthUser
	gofakeit.Struct(&user)
	userID, _ := suite.us.Create(user)

	var oldest uint
	var before uint
	var after uint
	var newest uint

	count := 7
	// - after - - before - -
	for i := 0; i < count; i++ {
		var u people.AuthUser
		gofakeit.Struct(&u)
		id, _ := suite.us.Create(u)
		suite.us.Follow(userID, u.Handle)
		var post people.PostBody
		gofakeit.Struct(&post)
		p, _ := suite.ps.Create(id, post)
		switch i {
		case 0:
			oldest = p.ID
		case 1:
			after = p.ID
		case 4:
			before = p.ID
		case count - 1:
			newest = p.ID
		}
	}

	tests := map[string]struct {
		pagination people.SeekPagination
		oldest     uint
		newest     uint
		count      int
	}{
		"no pagination":             {people.NewSeekPagination(nil, nil, nil), oldest, newest, count},
		"before":                    {people.NewSeekPagination(&before, nil, nil), oldest, before - 1, 4},
		"after":                     {people.NewSeekPagination(nil, &after, nil), after + 1, newest, 5},
		"after greater than before": {people.NewSeekPagination(&after, &before, nil), 0, 0, 0},
		"before and after":          {people.NewSeekPagination(&before, &after, nil), after + 1, before - 1, 2},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			res, _ := suite.ps.Feed(userID, tc.pagination)
			assert.Equal(suite.T(), tc.count, len(res.Data))
			if len(res.Data) == 0 {
				assert.Nil(suite.T(), res.Meta)
			} else {
				assert.Equal(suite.T(), tc.oldest, res.Meta.OldestID)
				assert.Equal(suite.T(), tc.newest, res.Meta.NewestID)
			}
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
