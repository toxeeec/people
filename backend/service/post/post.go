package post

import (
	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) people.PostService {
	return &service{db}
}

const (
	queryCreate = "INSERT INTO post(user_id, content) VALUES($1, $2) RETURNING post_id, content, created_at"
)

func (s *service) Create(userID uint, post people.PostBody) (people.Post, error) {
	var p people.Post
	return p, s.db.Get(&p, queryCreate, userID, post.Content)
}
