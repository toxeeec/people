package inmem

import (
	"errors"
	"fmt"
	"math"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type LikeKey struct {
	postID uint
	userID uint
}

type likeRepo struct {
	m      map[LikeKey]struct{}
	um     map[uint]people.User
	pm     map[uint]people.Post
	lastID uint
}

func NewLikeRepository(m map[LikeKey]struct{}, pm map[uint]people.Post, um map[uint]people.User) repository.Like {
	return &likeRepo{m: m, pm: pm, um: um}
}

func (r *likeRepo) Status(postID, userID uint) people.LikeStatus {
	_, ok := r.m[LikeKey{postID: postID, userID: userID}]
	return people.LikeStatus{IsLiked: ok}
}

func (r *likeRepo) Create(postID, userID uint) error {
	p, ok := r.pm[postID]
	if !ok {
		return fmt.Errorf("Like.Create: %w", repository.ErrPostNotFound)
	}
	p.Likes++

	key := LikeKey{postID: postID, userID: userID}
	_, ok = r.m[key]
	if ok {
		return fmt.Errorf("Like.Create: %w", repository.ErrAlreadyLiked)
	}
	r.m[key] = struct{}{}
	r.pm[postID] = p
	return nil
}

func (r *likeRepo) Delete(postID, userID uint) error {
	p, ok := r.pm[postID]
	if !ok {
		return fmt.Errorf("Like.Delete: %w", errors.New("Post not found"))
	}
	p.Likes--

	key := LikeKey{postID: postID, userID: userID}
	_, ok = r.m[key]
	if !ok {
		return fmt.Errorf("Like.Delete: %w", errors.New("Post not found"))
	}
	delete(r.m, key)
	r.pm[postID] = p
	return nil
}

func (r *likeRepo) ListPostLikes(postID uint, p *pagination.ID) ([]uint, error) {
	before := uint(math.MaxUint)
	if p.Before != nil {
		before = *p.Before
	}
	after := uint(0)
	if p.After != nil {
		after = *p.After
	}
	ids := []uint{}
	for k := range r.m {
		if k.postID == postID && k.userID < before && k.userID > after {
			ids = append(ids, k.userID)
		}
	}
	return ids, nil
}

func (r *likeRepo) ListUserLikes(userID uint, p *pagination.ID) ([]uint, error) {
	before := uint(math.MaxUint)
	after := uint(0)
	if p != nil {
		if p.Before != nil {
			before = *p.Before
		}
		if p.After != nil {
			after = *p.After
		}
	}
	ids := []uint{}
	for k := range r.m {
		if k.userID == userID && k.postID < before && k.postID > after {
			ids = append(ids, k.postID)
			if p != nil && len(ids) == int(p.Limit) {
				break
			}
		}
	}
	return ids, nil
}

func (r *likeRepo) ListStatusLiked(ids []uint, userID uint) (map[uint]struct{}, error) {
	lss := make(map[uint]struct{}, len(ids))
	for _, id := range ids {
		_, ok := r.m[LikeKey{postID: id, userID: userID}]
		if ok {
			lss[id] = struct{}{}
		}
	}
	return lss, nil
}

func (r *likeRepo) DeleteLike(ids []uint) error {
	for k, v := range r.pm {
		if contains(ids, k) {
			v.Likes--
			r.pm[k] = v
		}
	}
	return nil
}
