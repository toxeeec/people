package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	people "github.com/toxeeec/people/backend"
)

const day = 24 * time.Hour

const accessTokenDuration = 15 * time.Minute
const refreshTokenDuration = 30 * day

type refreshToken struct {
	ID     uuid.UUID `db:"token_id"`
	Value  string    `db:"value"`
	UserID uint      `db:"user_id"`
}

func newAccessToken(id uint) (string, error) {
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenDuration)),
		Subject:   fmt.Sprint(id),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	atString, err := at.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	return atString, err
}

const (
	queryInsertOrReplace = "INSERT INTO token(token_id, value, user_id) VALUES (:token_id, :value, :user_id) ON CONFLICT (token_id) DO UPDATE SET value = :value"
)

// newRefreshToken creates new uuid if nil is passed
func newRefreshToken(id uint, u *uuid.UUID) (refreshToken, error) {
	if u == nil {
		newUUID := uuid.New()
		u = &newUUID
	}
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenDuration)),
		Subject:   fmt.Sprint(id),
		ID:        u.String(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rtString, err := rt.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return refreshToken{}, err
	}

	return refreshToken{ID: *u, Value: rtString, UserID: id}, nil
}

func (s *service) NewTokens(id uint) (people.Tokens, error) {
	at, err := newAccessToken(id)
	if err != nil {
		return people.Tokens{}, err
	}

	rt, err := newRefreshToken(id, nil)
	if err != nil {
		return people.Tokens{}, err
	}

	_, err = s.db.NamedExec(queryInsertOrReplace, rt)
	if err != nil {
		return people.Tokens{}, err
	}

	return people.Tokens{AccessToken: at, RefreshToken: rt.Value}, nil
}

// VerifyCredentials returns id of the user.
func (s *service) VerifyCredentials(u people.AuthUser) (uint, error) {
	expected, err := s.us.Get(u.Handle)
	if err != nil {
		return 0, err
	}

	if err := u.Password.Compare(expected.Hash); err != nil {
		return 0, err
	}

	return expected.ID, nil
}
