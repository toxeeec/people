package auth

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service"
)

var (
	ErrInvalidSigningMethod = errors.New("Invalid signing method")
)

const (
	day = 24 * time.Hour

	// accessTokenDuration  = 15 * time.Minute
	accessTokenDuration  = 30 * day
	refreshTokenDuration = 30 * day
)

var (
	accessTokenSecret  = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	refreshTokenSecret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
)

// validateAccessToken returns user id if the tokenString is valid.
func ValidateAccessToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return accessTokenSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		sub := claims["sub"].(string)
		id, err := strconv.ParseUint(sub, 10, 64)
		return uint(id), err
	}
	return 0, errors.New("Unknown error")
}

func (s *authService) newTokens(id uint) (people.Tokens, error) {
	at, err := NewAccessToken(id)
	if err != nil {
		return people.Tokens{}, err
	}
	rt, err := NewRefreshToken(id, nil)
	if err != nil {
		return people.Tokens{}, err
	}
	err = s.tr.Create(rt)
	if err != nil {
		return people.Tokens{}, err
	}
	return people.Tokens{AccessToken: at, RefreshToken: rt.Value}, nil
}

func (s *authService) checkRefreshToken(rtString string) (people.RefreshToken, error) {
	rt, err := parseRefreshToken(rtString)
	if err != nil {
		return people.RefreshToken{}, service.NewError(people.AuthError, "Malformed refresh token")
	}
	if _, err := s.tr.Get(rt.Value); err != nil {
		// token doesn't exist
		go s.tr.Delete(rt.ID)
		return people.RefreshToken{}, service.NewError(people.AuthError, "Invalid refresh token")
	}

	return rt, nil
}

func NewAccessToken(userID uint) (string, error) {
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenDuration)),
		Subject:   fmt.Sprint(userID),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	atString, err := at.SignedString(accessTokenSecret)
	return atString, err
}

// NewRefreshToken creates new uuid if nil is passed.
func NewRefreshToken(userID uint, id *uuid.UUID) (people.RefreshToken, error) {
	if id == nil {
		newID := uuid.New()
		id = &newID
	}
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenDuration)),
		Subject:   fmt.Sprint(userID),
		ID:        id.String(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rtString, err := rt.SignedString(refreshTokenSecret)
	if err != nil {
		return people.RefreshToken{}, err
	}
	return people.RefreshToken{ID: *id, Value: rtString, UserID: userID}, nil
}

func parseRefreshToken(tokenString string) (people.RefreshToken, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return refreshTokenSecret, nil
	})
	if err != nil || !token.Valid {
		return people.RefreshToken{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		uuid, err := uuid.Parse(claims["jti"].(string))
		if err != nil {
			return people.RefreshToken{}, err
		}
		sub := claims["sub"].(string)
		userID, err := strconv.ParseUint(sub, 10, 64)
		if err != nil {
			return people.RefreshToken{}, err
		}
		return people.RefreshToken{ID: uuid, Value: tokenString, UserID: uint(userID)}, nil
	}
	return people.RefreshToken{}, errors.New("Unknown error")
}
