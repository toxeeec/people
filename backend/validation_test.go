package people_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
)

func TestUserValidate(t *testing.T) {
	t.Parallel()

	specialCharacters := people.AuthUser{Handle: "handle_"}
	valid := people.AuthUser{Handle: gofakeit.LetterN(10)}
	assert.Error(t, specialCharacters.Validate())
	assert.NoError(t, valid.Validate())
}

func TestPostBodyValidate(t *testing.T) {
	t.Parallel()

	empty := people.PostBody{Content: "\t\n \n\t"}
	empty.TrimContent()
	valid := people.PostBody{Content: gofakeit.Sentence(5)}
	valid.TrimContent()
	assert.Error(t, empty.Validate())
	assert.NoError(t, valid.Validate())
}
