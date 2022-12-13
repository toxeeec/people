package inmem_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestInmemTokenSuite(t *testing.T) {
	tm := map[uuid.UUID]people.RefreshToken{}
	um := map[uint]people.User{}
	tr := inmem.NewTokenRepository(tm)
	ur := inmem.NewUserRepository(um)
	fns := repotest.TestFns{SetupTest: func() {
		for k := range tm {
			delete(tm, k)
		}
		for k := range um {
			delete(um, k)
		}
	}}
	suite.Run(t, repotest.NewTokenSuite(tr, ur, fns))
}
