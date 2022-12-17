package repository

import (
	"errors"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
)

type Like interface {
	Status(postID, userID uint) people.LikeStatus
	Create(postID, userID uint) error
	Delete(postID, userID uint) error
	ListPostLikes(postID uint, p pagination.ID) (people.Users, error)
	ListStatusLiked(ids []uint, userID uint) (map[uint]struct{}, error)
	ListUserLikes(userID uint, p pagination.ID) ([]people.Post, error)
}

var (
	ErrPostNotFound = errors.New("Post not found")
	ErrAlreadyLiked = errors.New("This post is already liked by the current user")
)
