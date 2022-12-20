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
}
