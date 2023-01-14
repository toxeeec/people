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

const (
	PostFields = "post.post_id, post.user_id, content, replies_to, replies, likes, created_at"
	SelectPost = "SELECT " + PostFields + " FROM post"
)

func (r *postRepo) Create(np people.NewPost, userID uint, repliesTo *uint) (people.Post, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return people.Post{}, fmt.Errorf("Post.Create: %w", err)
	}
	defer tx.Rollback()
	const query = "INSERT INTO post(user_id, content, replies_to) VALUES ($1, $2, $3) RETURNING " + PostFields
	var p people.Post
	if err = r.db.Get(&p, query, userID, np.Content, repliesTo); err != nil {
		return p, fmt.Errorf("Post.Create: %w", err)
	}
	if repliesTo != nil {
		const incrementReplies = "UPDATE post SET replies = replies + 1 WHERE post_id = $1"
		_, err = tx.Exec(incrementReplies, *repliesTo)
		if err != nil {
			return people.Post{}, fmt.Errorf("Post.Create: %w", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return people.Post{}, fmt.Errorf("Post.Create: %w", err)
	}
	return p, nil
}

func (r *postRepo) Get(postID uint) (people.Post, error) {
	const query = SelectPost + " WHERE post_id = $1"
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

func (r *postRepo) List(ids []uint) ([]people.Post, error) {
	q, args, err := NewQuery(SelectPost).
		Where("post_id IN (?)", ids).
		Build()
	if err != nil {
		return nil, fmt.Errorf("Post.List: %w", err)
	}
	ps := []people.Post{}
	if err := r.db.Select(&ps, q, args...); err != nil {
		return nil, fmt.Errorf("Post.List: %w", err)
	}
	return ps, nil
}

func (r *postRepo) ListUserPosts(userID uint, p pagination.ID) ([]people.Post, error) {
	q, args, err := NewQuery(SelectPost).
		Where("user_id = ?", userID).
		Where("replies_to IS NULL").
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
	q, args, err := NewQuery(SelectPost).
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
	q, args, err := NewQuery(SelectPost).
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

func (r *postRepo) ListMatches(query string, p pagination.ID) ([]people.Post, error) {
	q, args, err := NewQuery(SelectPost).
		Where("ts @@ websearch_to_tsquery('english', ?)", query).
		Paginate(p, "post_id", "?").
		Build()
	if err != nil {
		return nil, fmt.Errorf("Post.ListMatches: %w", err)
	}
	ps := make([]people.Post, p.Limit)
	if err := r.db.Select(&ps, q, args...); err != nil {
		return nil, fmt.Errorf("Post.ListMatches: %w", err)
	}
	return ps, nil
}
