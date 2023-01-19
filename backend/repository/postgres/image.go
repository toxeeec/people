package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type imageRepo struct {
	db *sqlx.DB
}

func NewImageRepository(db *sqlx.DB) repository.Image {
	return &imageRepo{db}
}

func (r *imageRepo) Create(name string, userID uint) (people.ImageResponse, error) {
	const query = "INSERT INTO image(name, user_id) VALUES ($1, $2) RETURNING image_id"
	var ir people.ImageResponse
	err := r.db.Get(&ir, query, name, userID)
	if err != nil {
		return people.ImageResponse{}, fmt.Errorf("Image.Create: %w", err)
	}
	return ir, nil
}

func (r *imageRepo) Get(id uint) (people.Image, error) {
	const query = "SELECT * FROM image WHERE image_id = $1"
	var i people.Image
	err := r.db.Get(&i, query, id)
	if err != nil {
		return people.Image{}, fmt.Errorf("Image.Get: %w", err)
	}
	return i, nil
}

func (r *imageRepo) ListUnusedBefore(t time.Time) ([]people.Image, error) {
	const query = "SELECT * FROM image WHERE in_use = FALSE AND created_at < $1"
	var imgs []people.Image
	err := r.db.Select(&imgs, query, t)
	if err != nil {
		return nil, fmt.Errorf("Image.ListUnusedBefore: %w", err)
	}
	return imgs, nil
}

func (r *imageRepo) DeleteMany(ids []uint) {
	if len(ids) == 0 {
		return
	}
	q, args, err := NewQuery("DELETE FROM image").Where("image_id IN (?)", ids).Build()
	if err != nil {
		return
	}
	r.db.Exec(q, args...)
}

func (r *imageRepo) List(ids []uint) ([]people.Image, error) {
	if len(ids) == 0 {
		return []people.Image{}, nil
	}
	q, args, err := NewQuery("SELECT * FROM image").Where("image_id IN (?)", ids).Build()
	if err != nil {
		return nil, fmt.Errorf("Image.List: %w", err)
	}
	var imgs []people.Image
	err = r.db.Select(&imgs, q, args...)
	if err != nil {
		return nil, fmt.Errorf("Image.List: %w", err)
	}
	return imgs, nil
}

func (r *imageRepo) CreatePostImages(ids []uint, postID uint) error {
	if len(ids) == 0 {
		return nil
	}
	values := make([]string, 0, len(ids))
	for _, id := range ids {
		values = append(values, fmt.Sprintf("(%v, %v)", postID, id))
	}
	q, args, err := NewQuery("INSERT INTO post_image(post_id, image_id)").Values(values...).Build()
	if err != nil {
		return fmt.Errorf("Image.CreatePostImages: %w", err)
	}
	_, err = r.db.Exec(q, args...)
	if err != nil {
		return fmt.Errorf("Image.CreatePostImages: %w", err)
	}
	return nil
}

func (r *imageRepo) ListPostImages(postID uint) ([]people.Image, error) {
	const query = "SELECT * FROM image WHERE image_id IN (SELECT image_id FROM post_image WHERE post_id = $1)"
	var imgs []people.Image
	err := r.db.Select(&imgs, query, postID)
	if err != nil {
		return nil, fmt.Errorf("Image.ListPostImages: %w", err)
	}
	return imgs, nil
}

func (r *imageRepo) MarkUsed(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	q, args, err := NewQuery("UPDATE image SET in_use = TRUE").Where("image_id IN (?)", ids).Build()
	if err != nil {
		return fmt.Errorf("Image.MarkUsed: %w", err)
	}
	_, err = r.db.Exec(q, args...)
	if err != nil {
		return fmt.Errorf("Image.MarkUsed: %w", err)
	}
	return nil
}

func (r *imageRepo) DeleteManyPostImages(ids []uint) {
	if len(ids) == 0 {
		return
	}
	q, args, err := NewQuery("DELETE FROM post_image").Where("image_id IN (?)", ids).Build()
	if err != nil {
		return
	}
	r.db.Exec(q, args...)
}

type postImage struct {
	PostID  uint `db:"post_id"`
	ImageID uint `db:"image_id"`
}

func (r *imageRepo) ListPostsImageIDs(postIDs []uint) (map[uint][]uint, error) {
	if len(postIDs) == 0 {
		return map[uint][]uint{}, nil
	}
	q, args, err := NewQuery("SELECT post_id, image_id FROM post_image").Where("post_id IN (?)", postIDs).Build()
	if err != nil {
		return nil, fmt.Errorf("Image.ListPostsImages: %w", err)
	}
	var pimgs []postImage
	err = r.db.Select(&pimgs, q, args...)
	if err != nil {
		return nil, fmt.Errorf("Image.ListPostsImages: %w", err)
	}
	m := make(map[uint][]uint, len(postIDs))
	for _, pimg := range pimgs {
		postImgs := m[pimg.PostID]
		postImgs = append(postImgs, pimg.ImageID)
		m[pimg.PostID] = postImgs
	}
	return m, nil
}
