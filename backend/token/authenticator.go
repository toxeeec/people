package token

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	people "github.com/toxeeec/people/backend"
)

var (
	ErrNoAuthHeader      = errors.New("Authorization header is missing")
	ErrInvalidAuthHeader = errors.New("Authorization header is malformed")
)

func NewAuthenticator() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
		if ai.SecuritySchemeName != "bearerAuth" {
			return fmt.Errorf("security scheme %s != 'bearerAuth'", ai.SecuritySchemeName)
		}
		return authenticate(ai.RequestValidationInput.Request)
	}
}

func getJWTFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}

	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", ErrInvalidAuthHeader
	}

	return strings.TrimPrefix(authHeader, prefix), nil
}

func authenticate(r *http.Request) error {
	jwt, err := getJWTFromRequest(r)
	if err != nil {
		return err
	}

	id, err := ValidateAccessToken(jwt)
	if err != nil {
		return err
	}

	ctx := people.NewContext(r.Context(), people.UserIDKey, id)
	*r = *r.WithContext(ctx)

	return nil
}
