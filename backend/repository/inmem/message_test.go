package inmem_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestInmemMessageSuite(t *testing.T) {
	mm := make(map[uint]people.DBMessage)
	um := make(map[uint]people.User)
	mr := inmem.NewMessageRepository(mm)
	ur := inmem.NewUserRepository(um)
	fns := repotest.TestFns{SetupTest: func() {
		for k := range mm {
			delete(mm, k)
		}
		for k := range um {
			delete(um, k)
		}
	}}
	suite.Run(t, repotest.NewMessageSuite(mr, ur, fns))
}
