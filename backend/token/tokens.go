package token

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	people "github.com/toxeeec/people/backend"
)

var (
	ErrUnknown              = errors.New("Unknown error")
	ErrInvalidSigningMethod = errors.New("Invalid signing method")
)

const (
	day = 24 * time.Hour

	// accessTokenDuration  = 15 * time.Minute
	refreshTokenDuration = 30 * day

	accessTokenDuration = 1 * time.Minute
)

var (
	accessTokenSecret  = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	refreshTokenSecret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
)

func NewAccessToken(id uint) (string, error) {
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenDuration)),
		Subject:   fmt.Sprint(id),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	atString, err := at.SignedString(accessTokenSecret)
	return atString, err
}

// NewRefreshToken creates new uuid if nil is passed.
func NewRefreshToken(id uint, u *uuid.UUID) (people.RefreshToken, error) {
	if u == nil {
		id := uuid.New()
		u = &id
	}
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenDuration)),
		Subject:   fmt.Sprint(id),
		ID:        u.String(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rtString, err := rt.SignedString(refreshTokenSecret)
	if err != nil {
		return people.RefreshToken{}, err
	}

	return people.RefreshToken{ID: *u, Value: rtString, UserID: id}, nil
}

func ParseRefreshToken(tokenString string) (people.RefreshToken, error) {
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

	return people.RefreshToken{}, ErrUnknown
}

// ValidateAccessToken returns user id if the tokenString is valid.
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

	return 0, ErrUnknown
}
