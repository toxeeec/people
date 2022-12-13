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

func (r *likeRepo) ListUsers(postID uint, p pagination.ID) (people.Users, error) {
	before := uint(math.MaxUint)
	if p.Before != nil {
		before = *p.Before
	}
	after := uint(0)
	if p.Before != nil {
		after = *p.After
	}
	us := make([]people.User, 0, p.Limit)
	for k := range r.m {
		if k.postID == postID && k.userID < before && k.userID > after {
			us = append(us, r.um[k.userID])
		}
	}
	return pagination.NewResults[people.User, string](us), nil
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
