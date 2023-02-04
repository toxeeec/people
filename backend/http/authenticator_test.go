package http

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/toxeeec/people/backend/service/auth"
)

func newTokenWithNoneMethod(id uint) (string, error) {
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		Subject:   fmt.Sprint(id),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func TestAuthenticate(t *testing.T) {
	t.Parallel()

	expectedID := uint(rand.Uint32())
	validToken, _ := auth.NewAccessToken(expectedID)
	malformedToken := validToken[1:]

	noSigningMethod, _ := newTokenWithNoneMethod(expectedID)

	tests := map[string]struct {
		bearer string
		query  string
		valid  bool
	}{
		"no token":          {"", "", false},
		"no signing method": {"Bearer " + noSigningMethod, "", false},
		"malformed token":   {"Bearer " + string(malformedToken), "", false},
		"valid(header)":     {"Bearer " + validToken, "", true},
		"valid(query)":      {"", validToken, true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := http.Request{Header: http.Header{}, URL: &url.URL{}}
			if tc.bearer != "" {
				r.Header.Add("Authorization", tc.bearer)
			}
			if tc.query != "" {
				r.URL.RawQuery = fmt.Sprintf("access_token=%s", tc.query)
			}
			err := authenticate(&r)
			assert.Equal(t, tc.valid, err == nil)
		})
	}
}
