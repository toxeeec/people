package inmem

import (
	"errors"
	"fmt"
	"sort"
	"time"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type postImage struct {
	postID  uint
	imageID uint
}

type imageRepo struct {
	m      map[uint]people.Image
	pim    map[postImage]struct{}
	lastID uint
}

func NewImageRepository(m map[uint]people.Image) repository.Image {
	return &imageRepo{m: m, pim: map[postImage]struct{}{}}
}

func (r *imageRepo) newID() uint {
	r.lastID++
	return r.lastID
}

func (r *imageRepo) Create(path string, userID uint) (people.ImageResponse, error) {
	id := r.newID()
	r.m[id] = people.Image{
		ID:        id,
		Name:      path,
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
	imgs := []people.Image{}
	for _, v := range r.m {
		if v.InUse == false && v.CreatedAt.Before(t) {
			imgs = append(imgs, v)
		}
	}
	return imgs, nil
}

func (r *imageRepo) DeleteMany(ids []uint) {
	for _, id := range ids {
		delete(r.m, id)
	}
}

func (r *imageRepo) List(ids []uint) ([]people.Image, error) {
	if len(ids) == 0 {
		return []people.Image{}, nil
	}
	imgs := make([]people.Image, 0, len(ids))
	for _, img := range r.m {
		if contains(ids, img.ID) {
			imgs = append(imgs, img)
		}
	}
	return imgs, nil
}

func (r *imageRepo) CreatePostImages(ids []uint, postID uint) error {
	if len(ids) == 0 {
		return nil
	}
	for _, id := range ids {
		r.pim[postImage{postID: postID, imageID: id}] = struct{}{}
	}
	return nil
}

func (r *imageRepo) ListPostImages(postID uint) ([]people.Image, error) {
	imgs := make([]people.Image, 0, 4)
	for k := range r.pim {
		if k.postID == postID {
			imgs = append(imgs, r.m[k.imageID])
		}
	}
	return imgs, nil
}

func (r *imageRepo) MarkUsed(ids []uint) error {
	for k := range r.m {
		if contains(ids, k) {
			img := r.m[k]
			img.InUse = true
			r.m[k] = img
		}
	}
	return nil
}

func (r *imageRepo) DeleteManyPostImages(ids []uint) {
	for k := range r.pim {
		if contains(ids, k.imageID) {
			delete(r.pim, k)
		}
	}
}

func (r *imageRepo) ListPostsImageIDs(postIDs []uint) (map[uint][]uint, error) {
	m := make(map[uint][]uint, len(postIDs))
	for k := range r.pim {
		if contains(postIDs, k.postID) {
			imgs := m[k.postID]
			imgs = append(imgs, k.imageID)
			sort.Slice(imgs, func(i, j int) bool {
				return imgs[i] < imgs[j]
			})
			m[k.postID] = imgs
		}
	}
	return m, nil
}
