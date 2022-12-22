package inmem

import (
	"errors"
	"fmt"
	"strings"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
)

type userRepo struct {
	m      map[uint]people.User
	hashes map[uint]string
	lastID uint
}

func NewUserRepository(m map[uint]people.User) repository.User {
	return &userRepo{m: m, hashes: map[uint]string{}}
}

func (r *userRepo) newID() uint {
	r.lastID++
	return r.lastID
}

func (r *userRepo) GetID(handle string) (uint, error) {
	for _, v := range r.m {
		if v.Handle == handle {
			return v.ID, nil
		}
	}
	return 0, fmt.Errorf("User.GetID: %w", errors.New("User not found"))
}

func (r *userRepo) Create(au people.AuthUser) (people.User, error) {
	id := r.newID()
	r.m[id] = people.User{Handle: au.Handle, ID: id}
	r.hashes[id] = au.Password
	return r.m[id], nil
}

func (r *userRepo) Delete(id uint) error {
	delete(r.m, id)
	return nil
}

func (r *userRepo) GetHash(id uint) (string, error) {
	h, ok := r.hashes[id]
	if !ok {
		return "", fmt.Errorf("User.GetHash: %w", errors.New("User not found"))
	}
	return h, nil
}

func (r *userRepo) Get(id uint) (people.User, error) {
	u, ok := r.m[id]
	if !ok {
		return people.User{}, fmt.Errorf("User.Get: %w", errors.New("User not found"))
	}
	return u, nil
}

func (r *userRepo) List(ids []uint) ([]people.User, error) {
	us := make([]people.User, len(ids))
	for i, u := range r.m {
		if contains(ids, i) {
			us = append(us, u)
		}
	}
	return us, nil
}

func (r *userRepo) ListMatches(query string, p pagination.ID) ([]people.User, error) {
	us := make([]people.User, 0, p.Limit)
	for _, v := range r.m {
		if strings.Contains(strings.ToLower(v.Handle), strings.ToLower(query)) {
			us = append(us, v)
			if len(us) == int(p.Limit) {
				break
			}
		}
	}
	return us, nil
}
