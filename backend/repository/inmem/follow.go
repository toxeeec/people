package inmem

import (
	"errors"
	"fmt"
	"time"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type followRepo struct {
	m      map[FollowKey]time.Time
	um     map[uint]people.User
	lastID uint
}

type FollowKey struct {
	userID     uint
	followerID uint
}

func NewFollowRepository(m map[FollowKey]time.Time, us map[uint]people.User) repository.Follow {
	return &followRepo{m: m, um: us}
}

func (r *followRepo) GetStatusFollowing(srcID, userID uint) bool {
	_, ok := r.m[FollowKey{userID: userID, followerID: srcID}]
	return ok
}

func (r *followRepo) GetStatusFollowed(srcID, userID uint) bool {
	_, ok := r.m[FollowKey{userID: srcID, followerID: userID}]
	return ok
}

func (r *followRepo) Create(targetID, id uint) error {
	if targetID == id {
		return fmt.Errorf("Follow.Create: %w", repository.ErrSameUser)
	}
	fk := FollowKey{userID: targetID, followerID: id}
	_, ok := r.m[fk]
	if ok {
		return fmt.Errorf("Follow.Create: %w", repository.ErrAlreadyFollowed)
	}
	r.m[fk] = time.Now()

	target := r.um[targetID]
	target.Followers++
	r.um[targetID] = target
	u := r.um[id]
	u.Following++
	r.um[id] = u
	return nil
}

func (r *followRepo) Delete(targetID, id uint) error {
	fk := FollowKey{userID: targetID, followerID: id}
	_, ok := r.m[fk]
	if !ok {
		return fmt.Errorf("Follow.Delete: %w", errors.New("User not found"))
	}
	delete(r.m, fk)

	target := r.um[targetID]
	target.Followers--
	r.um[targetID] = target
	u := r.um[id]
	u.Following--
	r.um[id] = u
	return nil
}

func (r *followRepo) ListStatusFollowing(srcIDs []uint, userID uint) (map[uint]struct{}, error) {
	fss := make(map[uint]struct{}, len(srcIDs))
	for _, srcID := range srcIDs {
		_, ok := r.m[FollowKey{userID: userID, followerID: srcID}]
		if ok {
			fss[srcID] = struct{}{}
		}
	}
	return fss, nil
}

func (r *followRepo) ListStatusFollowed(srcIDs []uint, userID uint) (map[uint]struct{}, error) {
	fss := make(map[uint]struct{}, len(srcIDs))
	for _, srcID := range srcIDs {
		_, ok := r.m[FollowKey{userID: srcID, followerID: userID}]
		if ok {
			fss[srcID] = struct{}{}
		}
	}
	return fss, nil
}

func (r *followRepo) ListFollowing(id uint, p pagination.ID) ([]people.User, error) {
	before := time.Now()
	if p.Before != nil {
		before = r.m[FollowKey{userID: *p.Before, followerID: id}]
	}
	after := time.Time{}
	if p.After != nil {
		after = r.m[FollowKey{userID: *p.Before, followerID: id}]
	}
	var us []people.User
	for k, v := range r.m {
		if k.followerID == id && v.Before(before) && v.After(after) {
			us = append(us, r.um[k.userID])
		}
		if len(us) == int(p.Limit) {
			break
		}
	}
	return us, nil
}

func (r *followRepo) ListFollowers(id uint, p pagination.ID) ([]people.User, error) {
	before := time.Now()
	if p.Before != nil {
		before = r.m[FollowKey{userID: id, followerID: *p.Before}]
	}
	after := time.Time{}
	if p.After != nil {
		after = r.m[FollowKey{userID: id, followerID: *p.After}]
	}
	var us []people.User
	for k, v := range r.m {
		if k.userID == id && v.Before(before) && v.After(after) {
			us = append(us, r.um[k.followerID])
		}
		if len(us) == int(p.Limit) {
			break
		}
	}
	return us, nil
}
