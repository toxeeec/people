package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type tokenRepo struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) repository.Token {
	return &tokenRepo{db}
}

func (r *tokenRepo) Create(rt people.RefreshToken) error {
	const query = "INSERT INTO token VALUES (:token_id, :value, :user_id)"
	if _, err := r.db.NamedExec(query, rt); err != nil {
		return fmt.Errorf("Token.Create: %w", err)
	}
	return nil
}

func (r *tokenRepo) Get(value string) (people.RefreshToken, error) {
	const query = "SELECT token_id, value, user_id FROM token WHERE value = $1"
	var rt people.RefreshToken
	if err := r.db.Get(&rt, query, value); err != nil {
		return rt, fmt.Errorf("Token.Get: %w", err)
	}
	return rt, nil
}

func (r *tokenRepo) Delete(id uuid.UUID) error {
	const query = "DELETE FROM token WHERE token_id = $1"
	if _, err := r.db.Exec(query, id); err != nil {
		return fmt.Errorf("Token.Delete: %w", err)
	}
	return nil
}

func (r *tokenRepo) Update(rt people.RefreshToken) error {
	const query = "UPDATE token SET value = :value WHERE token_id = :token_id AND user_id = :user_id"
	if _, err := r.db.NamedExec(query, rt); err != nil {
		return fmt.Errorf("Token.Update: %w", err)
	}
	return nil
}
