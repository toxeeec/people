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

func (r *followRepo) ListStatusFollowing(userIDs []uint, srcID uint) (map[uint]struct{}, error) {
	fss := make(map[uint]struct{}, len(userIDs))
	for _, userID := range userIDs {
		_, ok := r.m[FollowKey{userID: srcID, followerID: userID}]
		if ok {
			fss[userID] = struct{}{}
		}
	}
	return fss, nil
}

func (r *followRepo) ListStatusFollowed(userIDs []uint, srcID uint) (map[uint]struct{}, error) {
	fss := make(map[uint]struct{}, len(userIDs))
	for _, userID := range userIDs {
		_, ok := r.m[FollowKey{userID: userID, followerID: srcID}]
		if ok {
			fss[userID] = struct{}{}
		}
	}
	return fss, nil
}

func (r *followRepo) ListFollowing(id uint, p *pagination.ID) ([]uint, error) {
	before := time.Now()
	after := time.Time{}
	if p != nil {
		if p.Before != nil {
			before = r.m[FollowKey{userID: *p.Before, followerID: id}]
		}
		if p.After != nil {
			after = r.m[FollowKey{userID: *p.Before, followerID: id}]
		}
	}
	var ids []uint
	for k, v := range r.m {
		if k.followerID == id && v.Before(before) && v.After(after) {
			ids = append(ids, k.userID)
		}
		if p != nil && len(ids) == int(p.Limit) {
			break
		}
	}
	return ids, nil
}

func (r *followRepo) ListFollowers(id uint, p *pagination.ID) ([]uint, error) {
	before := time.Now()
	after := time.Time{}
	if p != nil {
		if p.Before != nil {
			before = r.m[FollowKey{userID: *p.Before, followerID: id}]
		}
		if p.After != nil {
			after = r.m[FollowKey{userID: *p.Before, followerID: id}]
		}
	}
	var ids []uint
	for k, v := range r.m {
		if k.userID == id && v.Before(before) && v.After(after) {
			ids = append(ids, k.followerID)
		}
		if p != nil && len(ids) == int(p.Limit) {
			break
		}
	}
	return ids, nil
}

func (r *followRepo) DeleteFollower(userIDs []uint) error {
	for k, v := range r.um {
		if contains(userIDs, k) {
			v.Followers--
			r.um[k] = v
		}
	}
	return nil
}

func (r *followRepo) DeleteFollowing(userIDs []uint) error {
	for k, v := range r.um {
		if contains(userIDs, k) {
			v.Following--
			r.um[k] = v
		}
	}
	return nil
}
