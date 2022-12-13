package pagination_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository/inmem"
)

func TestIDPagination(t *testing.T) {
	um := map[uint]people.User{}
	ur := inmem.NewUserRepository(um)
	var users [2]people.User
	for i := range users {
		var au people.AuthUser
		gofakeit.Struct(&au)
		users[i], _ = ur.Create(au)
	}

	limit := uint(1)
	hp := pagination.New(&users[0].Handle, &users[1].Handle, &limit)
	p, err := pagination.Handle(hp).IDPagination(context.Background(), ur.GetID)
	assert.NoError(t, err)
	assert.Equal(t, users[0].ID, *p.Before)
	assert.Equal(t, users[1].ID, *p.After)
	assert.Equal(t, limit, p.Limit)
}
