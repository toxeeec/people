package repository

import (
	"time"

	people "github.com/toxeeec/people/backend"
)

type Image interface {
	Create(path string, userID uint) (people.ImageResponse, error)
	Get(id uint) (people.Image, error)
	ListUnusedBefore(t time.Time) ([]people.Image, error)
	DeleteMany(ids []uint)
	List(ids []uint) ([]people.Image, error)
	CreatePostImages(ids []uint, postID uint) error
	ListPostImages(postID uint) ([]people.Image, error)
	MarkUsed(ids []uint) error
	DeleteManyPostImages(ids []uint)
	ListPostsImageIDs(postIDs []uint) (map[uint][]uint, error)
}
