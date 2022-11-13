package post

import (
	"github.com/stretchr/testify/assert"
)

func (suite *PostSuite) TestLike() {
	suite.ps.Like(suite.post1.ID, suite.user2ID)

	tests := map[string]struct {
		postID uint
		userID uint
		valid  bool
	}{
		"invalid post id": {suite.post1.ID + 5, suite.user1ID, false},
		"invalid user id": {suite.post1.ID, suite.user1ID + 5, false},
		"already liked":   {suite.post1.ID, suite.user2ID, false},
		"valid":           {suite.post2.ID, suite.user1ID, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			likes, err := suite.ps.Like(tc.postID, tc.userID)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), uint(1), likes.Likes)
				assert.True(suite.T(), likes.IsLiked)
				post, _ := suite.ps.Get(tc.postID, &tc.userID)
				assert.True(suite.T(), post.IsLiked)
			}
		})
	}
}

func (suite *PostSuite) TestUnlike() {
	suite.ps.Like(suite.post1.ID, suite.user1ID)

	tests := map[string]struct {
		postID uint
		userID uint
		valid  bool
	}{
		"invalid post id": {suite.post1.ID + 5, suite.user1ID, false},
		"invalid user id": {suite.post1.ID, suite.user1ID + 5, false},
		"not liked":       {suite.post1.ID, suite.user2ID, false},
		"valid":           {suite.post1.ID, suite.user1ID, true},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			likes, err := suite.ps.Unlike(tc.postID, tc.userID)
			assert.Equal(suite.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(suite.T(), uint(0), likes.Likes)
				assert.False(suite.T(), likes.IsLiked)
				post, _ := suite.ps.Get(tc.postID, &tc.userID)
				assert.False(suite.T(), post.IsLiked)
			}
		})
	}
}
