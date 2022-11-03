package post

import (
	"errors"

	"github.com/lib/pq"
	people "github.com/toxeeec/people/backend"
)

var (
	ErrInvalidPostID = errors.New("Invalid post id")
	ErrAlreadyLiked  = errors.New("This post is already liked by the current user")
)

const (
	queryLike           = "INSERT INTO post_like(post_id, user_id) VALUES($1, $2) RETURNING post_id"
	queryIncrementLikes = "UPDATE post SET likes = likes + 1 WHERE post_id = $1 RETURNING likes"
)

func (s *service) Like(postID, userID uint) (people.Likes, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return people.Likes{}, err
	}
	defer tx.Rollback()

	err = tx.Get(new(uint), queryLike, postID, userID)
	if err != nil {
		var e *pq.Error
		if errors.As(err, &e) {
			if e.Constraint == "post_like_pkey" {
				return people.Likes{}, ErrAlreadyLiked
			}
			if e.Constraint == "post_like_post_id_fkey" {
				return people.Likes{}, ErrInvalidPostID
			}
		}
		return people.Likes{}, err
	}

	var l people.Likes
	err = tx.Get(&l, queryIncrementLikes, postID)
	if err != nil {
		return people.Likes{}, err
	}

	return l, tx.Commit()
}
