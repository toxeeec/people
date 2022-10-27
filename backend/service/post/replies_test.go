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
