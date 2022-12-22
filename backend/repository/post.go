package repository

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
)

type Post interface {
	Create(np people.NewPost, userID uint, repliesTo *uint) (people.Post, error)
	Get(postID uint) (people.Post, error)
	Delete(postID, userID uint) error
	ListUserPosts(userID uint, p pagination.ID) ([]people.Post, error)
	ListFeed(followingIDs []uint, userID uint, p pagination.ID) ([]people.Post, error)
	ListReplies(postID uint, p pagination.ID) ([]people.Post, error)
	ListMatches(query string, p pagination.ID) ([]people.Post, error)
}
