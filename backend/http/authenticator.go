package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/toxeeec/people/backend/service/auth"
)

func (h *handler) newAuthenticator() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
		if ai.SecuritySchemeName != "bearerAuth" {
			return fmt.Errorf("security scheme %s != 'bearerAuth'", ai.SecuritySchemeName)
		}
		return authenticate(ai.RequestValidationInput.Request)
	}
}

func authenticate(r *http.Request) error {
	jwt, err := getJWTFromRequest(r)
	if err != nil {
		return err
	}
	id, err := auth.ValidateAccessToken(jwt)
	if err != nil {
		return err
	}
	ctx := newContext(r.Context(), userIDKey, id)
	*r = *r.WithContext(ctx)
	return nil
}

func getJWTFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header is missing")
	}
	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("Authorization header is malformed")
	}
	return strings.TrimPrefix(authHeader, prefix), nil
}
