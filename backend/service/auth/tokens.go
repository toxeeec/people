package auth

import (
	"github.com/google/uuid"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/token"
)

const (
	queryInsert = "INSERT INTO token(token_id, value, user_id) VALUES (:token_id, :value, :user_id)"
	queryUpdate = "UPDATE token SET value = $1 WHERE token_id = $2 AND user_id = $3"
	queryExists = "SELECT EXISTS(SELECT 1 FROM token WHERE value = $1)"
	queryDelete = "DELETE FROM token WHERE token_id = $1"
)

func (s *service) NewTokens(id uint) (people.Tokens, error) {
	at, err := token.NewAccessToken(id)
	if err != nil {
		return people.Tokens{}, err
	}

	rt, err := token.NewRefreshToken(id, nil)
	if err != nil {
		return people.Tokens{}, err
	}

	_, err = s.db.NamedExec(queryInsert, rt)
	if err != nil {
		return people.Tokens{}, err
	}

	return people.Tokens{AccessToken: &at, RefreshToken: rt.Value}, nil
}

func (s *service) UpdateRefreshToken(userID uint, tokenID uuid.UUID) (token.RefreshToken, error) {
	rt, err := token.NewRefreshToken(userID, &tokenID)
	if err != nil {
		return token.RefreshToken{}, err
	}

	_, err = s.db.Exec(queryUpdate, rt.Value, tokenID, userID)
	if err != nil {
		return token.RefreshToken{}, err
	}

	return rt, nil
}

func (s *service) CheckRefreshToken(token token.RefreshToken) bool {
	var exists bool
	s.db.Get(&exists, queryExists, token.Value)
	if exists {
		return true
	}

	s.db.Exec(queryDelete, token.ID)
	return false
}
