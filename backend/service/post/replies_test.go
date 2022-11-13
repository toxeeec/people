package post

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func (suite *PostSuite) TestCreateReply() {
	r, _ := suite.ps.CreateReply(suite.post1.ID, suite.user1ID, suite.replyBody)
	rows, err := suite.ps.db.Queryx(`SELECT post_id, content, replies_to, replies FROM post WHERE replies_to = $1`, suite.post1.ID)
	assert.NoError(suite.T(), err)
	for rows.Next() {
		var actual people.Post
		rows.StructScan(&actual)
		assert.Equal(suite.T(), r.ID, actual.ID)
		assert.Equal(suite.T(), r.Content, actual.Content)
		assert.Equal(suite.T(), suite.post1.ID, uint(actual.RepliesTo.Int32))

		p, _ := suite.ps.Get(suite.post1.ID, nil)
		assert.Equal(suite.T(), uint(1), p.Replies)
	}
}

func (suite *PostSuite) TestReplies() {
	count := 5
	for i := 0; i < count; i++ {
		var p people.PostBody
		gofakeit.Struct(&p)
		suite.ps.CreateReply(suite.post1.ID, suite.user1ID, p)
	}

	tests := map[string]struct {
		id       uint
		expected int
	}{
		"invalid id": {suite.post1.ID + 5, 0},
		"0 replies":  {suite.post2.ID, 0},
		"valid":      {suite.post1.ID, count},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			pagination := people.NewPagination[uint](nil, nil, nil)
			posts, _ := suite.ps.Replies(tc.id, nil, pagination)
			assert.Equal(suite.T(), tc.expected, len(posts.Data))
		})
	}
}
