package inmem

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type tokenRepo struct {
	m map[uuid.UUID]people.RefreshToken
}

func NewTokenRepository(m map[uuid.UUID]people.RefreshToken) repository.Token {
	return &tokenRepo{m}
}

func (r *tokenRepo) Create(rt people.RefreshToken) error {
	_, ok := r.m[rt.ID]
	if ok {
		return fmt.Errorf("Token.Create: %w", errors.New("Token with this id already exists"))
	}
	r.m[rt.ID] = rt
	return nil
}

func (r *tokenRepo) Get(value string) (people.RefreshToken, error) {
	for _, v := range r.m {
		if v.Value == value {
			return v, nil
		}
	}
	return people.RefreshToken{}, fmt.Errorf("Token.Get: %w", errors.New("Token not found"))
}

func (r *tokenRepo) Delete(id uuid.UUID) error {
	delete(r.m, id)
	return nil
}

func (r *tokenRepo) Update(rt people.RefreshToken) error {
	t := r.m[rt.ID]
	if t.UserID == rt.UserID {
		t.Value = rt.Value
		r.m[rt.ID] = t
	}
	return nil
}
