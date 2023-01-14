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
	//ListPostLikes returns ids of users that liked the post
	ListPostLikes(postID uint, p *pagination.ID) ([]uint, error)
	//ListUserLikes returns ids of posts liked by the user
	ListUserLikes(userID uint, p *pagination.ID) ([]uint, error)
	ListStatusLiked(ids []uint, userID uint) (map[uint]struct{}, error)
	DeleteLike(ids []uint) error
}

var (
	ErrPostNotFound = errors.New("Post not found")
	ErrAlreadyLiked = errors.New("This post is already liked by the current user")
)
