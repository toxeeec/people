package inmem_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestInmemFollowSuite(t *testing.T) {
	um := map[uint]people.User{}
	fm := map[inmem.FollowKey]time.Time{}
	ur := inmem.NewUserRepository(um)
	fr := inmem.NewFollowRepository(fm, um)
	fns := repotest.TestFns{SetupTest: func() {
		for k := range um {
			delete(um, k)
		}
		for k := range fm {
			delete(fm, k)
		}
	}}
	suite.Run(t, repotest.NewFollowSuite(fr, ur, fns))
}
