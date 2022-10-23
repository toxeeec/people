package people_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func TestUserValidate(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		user  people.AuthUser
		valid bool
	}{
		"special characters": {people.AuthUser{Handle: "handle_"}, false},
		"valid":              {people.AuthUser{Handle: gofakeit.LetterN(5)}, true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := tc.user.Validate()
			if tc.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
