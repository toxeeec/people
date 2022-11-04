package post

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *PostSuite) TestLike() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	var post1 people.PostBody
	var post2 people.PostBody
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	gofakeit.Struct(&post1)
	gofakeit.Struct(&post2)
	id1, _ := suite.us.Create(user1)
	id2, _ := suite.us.Create(user2)
	p1, _ := suite.ps.Create(id1, post1)

	suite.ps.db.MustExec("INSERT INTO post_like(post_id, user_id) VALUES ($1, $2)", p1.ID, id2)

	tests := map[string]struct {
		postID uint
		userID uint
		valid  bool
	}{
		"invalid post id": {p1.ID + 1, id1, false},
		"invalid user id": {p1.ID, id1 + 2, false},
		"already liked":   {p1.ID, id2, false},
		"valid":           {p1.ID, id1, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			likes, err := suite.ps.Like(tc.postID, tc.userID)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), uint(1), likes.Likes)
			}
		})
	}
}

func (suite *PostSuite) TestUnlike() {
	var user1 people.AuthUser
	var user2 people.AuthUser
	var post1 people.PostBody
	var post2 people.PostBody
	gofakeit.Struct(&user1)
	gofakeit.Struct(&user2)
	gofakeit.Struct(&post1)
	gofakeit.Struct(&post2)
	id1, _ := suite.us.Create(user1)
	id2, _ := suite.us.Create(user2)
	p1, _ := suite.ps.Create(id1, post1)

	suite.ps.Like(p1.ID, id1)

	tests := map[string]struct {
		postID uint
		userID uint
		valid  bool
	}{
		"invalid post id": {p1.ID + 1, id1, false},
		"invalid user id": {p1.ID, id1 + 2, false},
		"not liked":       {p1.ID, id2, false},
		"valid":           {p1.ID, id1, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			likes, err := suite.ps.Unlike(tc.postID, tc.userID)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), uint(0), likes.Likes)
			}
		})
	}
}
