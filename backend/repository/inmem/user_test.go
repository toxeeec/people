package inmem_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestInmemUserSuite(t *testing.T) {
	m := map[uint]people.User{}
	ur := inmem.NewUserRepository(m)
	fns := repotest.TestFns{SetupTest: func() {
		for k := range m {
			delete(m, k)
		}
	}}
	suite.Run(t, repotest.NewUserSuite(ur, fns))
}
