package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) repository.Post {
	return &postRepo{db}
}

func (r *postRepo) Create(np people.NewPost, userID uint, repliesTo *uint) (people.Post, error) {
	const query = "INSERT INTO post(user_id, content, replies_to) VALUES ($1, $2, $3) RETURNING *"
	var p people.Post
	if err := r.db.Get(&p, query, userID, np.Content, repliesTo); err != nil {
		return p, fmt.Errorf("Post.Create: %w", err)
	}
	return p, nil
}

func (r *postRepo) Get(postID uint) (people.Post, error) {
	const query = "SELECT * FROM post WHERE post_id = $1"
	var p people.Post
	if err := r.db.Get(&p, query, postID); err != nil {
		return p, fmt.Errorf("Post.Get: %w", err)
	}
	return p, nil
}

func (r *postRepo) Delete(postID, userID uint) error {
	const query = "DELETE FROM post WHERE post_id = $1 AND user_id = $2"
	if _, err := r.db.Exec(query, postID, userID); err != nil {
		return fmt.Errorf("Post.Delete: %w", err)
	}
	return nil
}

func (r *postRepo) ListUserPosts(userID uint, p pagination.ID) ([]people.Post, error) {
	q, args, err := NewQuery("SELECT * FROM post").
		Where("user_id = ?", userID).
		Paginate(p, "post_id", "?").
		Build()
	if err != nil {
		return nil, fmt.Errorf("Post.ListUserPosts: %w", err)
	}
	ps := make([]people.Post, p.Limit)
	if err := r.db.Select(&ps, q, args...); err != nil {
		return nil, fmt.Errorf("Post.ListUserPosts: %w", err)
	}
	return ps, nil
}

func (r *postRepo) ListFeed(followingIDs []uint, userID uint, p pagination.ID) ([]people.Post, error) {
	if len(followingIDs) == 0 {
		return []people.Post{}, nil
	}
	q, args, err := NewQuery("SELECT * FROM post").
		Where("post.user_id IN (?)", append(followingIDs, userID)).
		Where("replies_to IS NULL").
		Paginate(p, "post_id", "?").
		Build()
	if err != nil {
		return nil, fmt.Errorf("Post.ListFeed: %w", err)
	}
	ps := make([]people.Post, p.Limit)
	if err := r.db.Select(&ps, q, args...); err != nil {
		return nil, fmt.Errorf("Post.ListFeed: %w", err)
	}
	return ps, nil
}

func (r *postRepo) ListReplies(postID uint, p pagination.ID) ([]people.Post, error) {
	q, args, err := NewQuery("SELECT * FROM post").
		Where("replies_to = ?", postID).
		Paginate(p, "post_id", "?").
		Build()
	if err != nil {
		return nil, fmt.Errorf("Post.ListReplies: %w", err)
	}
	ps := make([]people.Post, p.Limit)
	if err := r.db.Select(&ps, q, args...); err != nil {
		return nil, fmt.Errorf("Post.ListReplies: %w", err)
	}
	return ps, nil
}
