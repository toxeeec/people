package inmem_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestInmemMessageSuite(t *testing.T) {
	msgs := make(map[uint][]people.DBMessage)
	threads := make(map[uint]struct{})
	threadUsers := make(map[uint][]uint)
	um := make(map[uint]people.User)
	mr := inmem.NewMessageRepository(msgs, threads, threadUsers, um)
	ur := inmem.NewUserRepository(um)
	fns := repotest.TestFns{SetupTest: func() {
		for k := range msgs {
			delete(msgs, k)
		}
		for k := range threads {
			delete(threads, k)
		}
		for k := range threadUsers {
			delete(threadUsers, k)
		}
		for k := range um {
			delete(um, k)
		}
	}}
	suite.Run(t, repotest.NewMessageSuite(mr, ur, fns))
}
