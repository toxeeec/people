package postgres

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type likeRepo struct {
	db *sqlx.DB
}

func NewLikeRepository(db *sqlx.DB) repository.Like {
	return &likeRepo{db}
}

func (r *likeRepo) Status(postID, userID uint) people.LikeStatus {
	const query = "SELECT EXISTS(SELECT 1 FROM post_like WHERE post_id = $1 AND user_id = $2) AS is_liked"
	var ls people.LikeStatus
	r.db.Get(&ls, query, postID, userID)
	return ls
}

func (r *likeRepo) Create(postID, userID uint) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("Like.Create: %w", err)
	}
	defer tx.Rollback()
	const query = "INSERT INTO post_like(post_id, user_id) VALUES ($1, $2) RETURNING post_id"
	err = tx.Get(new(uint), query, postID, userID)
	if err != nil {
		var e *pq.Error
		if errors.As(err, &e) {
			if e.Constraint == "post_like_pkey" {
				return fmt.Errorf("Like.Create: %w", repository.ErrAlreadyLiked)
			}
		}
		return fmt.Errorf("Like.Create: %w", repository.ErrPostNotFound)
	}

	const incrementLikes = "UPDATE post SET likes = likes + 1 WHERE post_id = $1"
	_, err = tx.Exec(incrementLikes, postID)
	if err != nil {
		return fmt.Errorf("Like.Create: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Like.Create: %w", err)
	}
	return nil
}

func (r *likeRepo) Delete(postID, userID uint) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("Like.Delete: %w", err)
	}
	defer tx.Rollback()
	const query = "DELETE FROM post_like WHERE post_id = $1 AND user_id = $2 RETURNING post_id"
	err = tx.Get(new(uint), query, postID, userID)
	if err != nil {
		return fmt.Errorf("Like.Delete: %w", err)
	}

	const decrementLikes = "UPDATE post SET likes = likes - 1 WHERE post_id = $1 RETURNING likes"
	_, err = tx.Exec(decrementLikes, postID)
	if err != nil {
		return fmt.Errorf("Like.Delete: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Like.Delete: %w", err)
	}
	return nil
}

func (r *likeRepo) ListPostLikes(postID uint, p pagination.ID) (people.Users, error) {
	const paginationValue = "(SELECT liked_at FROM post_like WHERE user_id = ? AND post_id = ?)"
	q, args, err := NewQuery(SelectUser).
		Join("post_like", "user_profile.user_id = post_like.user_id").
		Where("post_id = ?", postID).
		Paginate(p, "liked_at", paginationValue, postID).
		Build()
	if err != nil {
		return people.Users{}, fmt.Errorf("Like.ListUsers: %w", err)
	}
	us := make([]people.User, p.Limit)
	if err := r.db.Select(&us, q, args...); err != nil {
		return people.Users{}, fmt.Errorf("Like.ListUsers: %w", err)
	}
	return pagination.NewResults[people.User, string](us), nil
}

func (r *likeRepo) ListUserLikes(userID uint, p pagination.ID) ([]people.Post, error) {
	const paginationValue = "(SELECT liked_at FROM post_like WHERE user_id = ? AND post_id = ?)"
	q, args, err := NewQuery(SelectPost).
		Join("post_like", "post_like.post_id = post.post_id").
		Where("post_like.user_id = ?", userID).
		Paginate(p, "liked_at", paginationValue, userID).
		Build()
	if err != nil {
		return nil, fmt.Errorf("Post.ListUserLikes: %w", err)
	}
	ps := make([]people.Post, p.Limit)
	if err := r.db.Select(&ps, q, args...); err != nil {
		return nil, fmt.Errorf("Post.ListUserLikes: %w", err)
	}
	return ps, nil
}

func (r *likeRepo) ListStatusLiked(ids []uint, userID uint) (map[uint]struct{}, error) {
	if len(ids) == 0 {
		return make(map[uint]struct{}), nil
	}
	const query = "SELECT post_id FROM post_like"
	q, args, err := NewQuery(query).Where("post_id IN (?)", ids).Where("user_id = ?", userID).Build()
	if err != nil {
		return nil, fmt.Errorf("Like.ListStatusLiked: %w", err)
	}
	postIDs := make([]uint, len(ids))
	err = r.db.Select(&postIDs, q, args...)
	if err != nil {
		return nil, fmt.Errorf("Like.ListStatusLiked: %w", err)
	}
	lss := make(map[uint]struct{}, len(ids))
	for _, id := range postIDs {
		lss[id] = struct{}{}
	}
	return lss, nil
}
