package post

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *PostSuite) TestCreateReply() {
	var user people.AuthUser
	var post people.PostBody
	gofakeit.Struct(&user)
	gofakeit.Struct(&post)
	userID, _ := suite.us.Create(user)
	p, _ := suite.ps.Create(userID, post)
	postID := p.ID

	var expected people.PostBody
	gofakeit.Struct(&expected)
	p, _ = suite.ps.CreateReply(postID, userID, expected)

	rows, err := suite.ps.db.Queryx(`SELECT post_id, content, replies_to, replies AS "user.user_id" FROM post WHERE replies_to = $1`, postID)
	assert.NoError(suite.T(), err)
	for rows.Next() {
		var actual people.Post
		rows.StructScan(&actual)
		assert.Equal(suite.T(), p.ID, actual.ID)
		assert.Equal(suite.T(), expected.Content, actual.Content)
		assert.Equal(suite.T(), postID, uint(actual.RepliesTo.Int32))

		p, _ := suite.ps.Get(postID)
		assert.Equal(suite.T(), uint(1), p.Replies)
	}
}

func (suite *PostSuite) TestReplies() {
	var user people.AuthUser
	var post1 people.PostBody
	var post2 people.PostBody
	gofakeit.Struct(&user)
	gofakeit.Struct(&post1)
	gofakeit.Struct(&post2)
	userID, _ := suite.us.Create(user)
	p1, _ := suite.ps.Create(userID, post1)
	p2, _ := suite.ps.Create(userID, post2)

	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.CreateReply(p1.ID, userID, p)
	}

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"invalid id": {p1.ID + 2, 0},
		"0 replies":  {p2.ID, 0},
		"valid":      {p1.ID, count},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			pagination := people.NewSeekPagination(nil, nil, nil)
			posts, _ := suite.ps.Replies(tc.id, pagination)
			assert.Equal(suite.T(), tc.expected, len(posts.Data))
		})
	}
}
