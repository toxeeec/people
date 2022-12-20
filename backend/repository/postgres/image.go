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

func (r *imageRepo) Create(path string, userID uint) (people.ImageResponse, error) {
	const query = "INSERT INTO image(path, user_id) VALUES($1, $2) RETURNING image_id"
	var ir people.ImageResponse
	err := r.db.Get(&ir, query, path, userID)
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
	var is []people.Image
	err := r.db.Select(&is, query, t)
	if err != nil {
		return nil, fmt.Errorf("Image.ListUnusedBefore: %w", err)
	}
	return is, nil
}

func (r *imageRepo) DeleteMany(ids []uint) {
	const query = "DELETE FROM image WHERE image_id IN (?)"
	q, args, err := NewQuery("DELETE FROM image").Where("image_id IN (?)", ids).Build()
	if err != nil {
		return
	}
	r.db.Exec(q, args...)
}
