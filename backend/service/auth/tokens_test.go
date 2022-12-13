package auth

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAccessToken(t *testing.T) {
	t.Parallel()

	expected := uint(1)
	at, _ := NewAccessToken(expected)
	actual, err := ValidateAccessToken(at)
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)

	malformedToken := []byte(at)
	malformedToken[rand.Intn(len(malformedToken))] += 1

	_, err = ValidateAccessToken(string(malformedToken))
	assert.Error(t, err)
}

func TestParseRefreshToken(t *testing.T) {
	t.Parallel()

	expected, _ := NewRefreshToken(1, nil)
	actual, _ := parseRefreshToken(expected.Value)
	assert.Equal(t, expected.Value, actual.Value)
}
