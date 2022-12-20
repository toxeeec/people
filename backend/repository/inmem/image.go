package inmem

import (
	"errors"
	"fmt"
	"time"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type imageRepo struct {
	m      map[uint]people.Image
	lastID uint
}

func NewImageRepository(m map[uint]people.Image) repository.Image {
	return &imageRepo{m: m}
}

func (r *imageRepo) newID() uint {
	r.lastID++
	return r.lastID
}

func (r *imageRepo) Create(path string, userID uint) (people.ImageResponse, error) {
	id := r.newID()
	r.m[id] = people.Image{
		ID:        id,
		Path:      path,
		CreatedAt: time.Now(),
		UserID:    userID,
	}
	return people.ImageResponse{ID: id}, nil
}

func (r *imageRepo) Get(id uint) (people.Image, error) {
	i, ok := r.m[id]
	if !ok {
		return people.Image{}, fmt.Errorf("Image.Get: %w", errors.New("Image not found"))
	}
	return i, nil
}

func (r *imageRepo) ListUnusedBefore(t time.Time) ([]people.Image, error) {
	const query = "SELECT * FROM image WHERE in_use = FALSE AND created_at < $1"
	is := []people.Image{}
	for _, v := range r.m {
		if v.InUse == false && v.CreatedAt.Before(t) {
			is = append(is, v)
		}
	}
	return is, nil
}

func (r *imageRepo) DeleteMany(ids []uint) {
	for _, id := range ids {
		delete(r.m, id)
	}
}
