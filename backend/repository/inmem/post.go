package inmem

import (
	"errors"
	"fmt"
	"math"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type postRepo struct {
	m      map[uint]people.Post
	lastID uint
}

func NewPostRepository(m map[uint]people.Post) repository.Post {
	return &postRepo{m: m}
}

func (r *postRepo) newID() uint {
	r.lastID++
	return r.lastID
}

func (r *postRepo) Create(np people.NewPost, userID uint, repliesTo *uint) (people.Post, error) {
	if repliesTo != nil {
		p, ok := r.m[*repliesTo]
		if !ok {
			return people.Post{}, fmt.Errorf("Post.Get: %w", errors.New("Post not found"))
		}
		p.Replies++
		r.m[*repliesTo] = p
	}
	id := r.newID()
	r.m[id] = people.Post{ID: id, Content: np.Content, UserID: userID, RepliesTo: repliesTo}
	return r.m[id], nil
}

func (r *postRepo) Get(postID uint) (people.Post, error) {
	p, ok := r.m[postID]
	if !ok {
		return people.Post{}, fmt.Errorf("Post.Get: %w", errors.New("Post not found"))
	}
	return p, nil
}

func (r *postRepo) Delete(postID, userID uint) error {
	p, ok := r.m[postID]
	if ok && p.UserID == userID {
		delete(r.m, postID)
	}
	return nil
}

func (r *postRepo) ListUserPosts(userID uint, p pagination.Pagination[uint]) ([]people.Post, error) {
	before := uint(math.MaxUint)
	if p.Before != nil {
		before = *p.Before
	}
	after := uint(0)
	if p.Before != nil {
		after = *p.After
	}
	var ps []people.Post
	for k, v := range r.m {
		if v.UserID == userID && k < before && k > after {
			ps = append(ps, v)
		}
		if len(ps) == int(p.Limit) {
			break
		}
	}
	return ps, nil
}

func (r *postRepo) ListFeed(followingIDs []uint, userID uint, p pagination.ID) ([]people.Post, error) {
	before := uint(math.MaxUint)
	if p.Before != nil {
		before = *p.Before
	}
	after := uint(0)
	if p.Before != nil {
		after = *p.After
	}
	followingIDs = append(followingIDs, userID)
	var ps []people.Post
	for k, v := range r.m {
		if contains(followingIDs, v.UserID) && k < before && k > after && v.RepliesTo == nil {
			ps = append(ps, v)
		}
		if len(ps) == int(p.Limit) {
			break
		}
	}
	return ps, nil
}

func (r *postRepo) ListReplies(postID uint, p pagination.ID) ([]people.Post, error) {
	before := uint(math.MaxUint)
	if p.Before != nil {
		before = *p.Before
	}
	after := uint(0)
	if p.Before != nil {
		after = *p.After
	}
	ps := make([]people.Post, 0, p.Limit)
	for k, v := range r.m {
		if v.RepliesTo != nil && *v.RepliesTo == postID && k < before && k > after {
			ps = append(ps, v)
		}
	}
	return ps, nil
}

func contains(s []uint, item uint) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}
