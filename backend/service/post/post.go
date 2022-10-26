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
	queryGet    = `SELECT content, created_at, user_profile.handle AS "user.handle" FROM post JOIN user_profile ON post.user_id = user_profile.user_id WHERE post_id = $1`
	queryDelete = "DELETE FROM post WHERE post_id = $1 AND user_id = $2 RETURNING post_id"
)

func (s *service) Create(userID uint, post people.PostBody) (people.Post, error) {
	var p people.Post
	return p, s.db.Get(&p, queryCreate, userID, post.Content)
}

func (s *service) Get(id uint) (people.Post, error) {
	var p people.Post
	return p, s.db.Get(&p, queryGet, id)
}

func (s *service) Delete(postID, userID uint) error {
	return s.db.Get(new(uint), queryDelete, postID, userID)
}
