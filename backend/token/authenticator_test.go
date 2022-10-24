package token

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
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
	validToken, _ := NewAccessToken(expectedID)

	malformedToken := []rune(validToken)
	malformedToken[rand.Intn(len(malformedToken))] += 1

	noSigningMethod, _ := newTokenWithNoneMethod(expectedID)

	tests := map[string]struct {
		bearer string
		valid  bool
	}{
		"no token":          {"", false},
		"no signing method": {"Bearer " + noSigningMethod, false},
		"malformed token":   {"Bearer " + string(malformedToken), false},
		"valid":             {"Bearer " + validToken, true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := http.Request{}
			r.Header = make(http.Header)
			r.Header.Add("Authorization", tc.bearer)
			err := authenticate(&r)
			assert.Equal(t, tc.valid, err == nil)
		})
	}
}
