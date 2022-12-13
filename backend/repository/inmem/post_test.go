package inmem_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestInmemPostSuite(t *testing.T) {
	pm := map[uint]people.Post{}
	um := map[uint]people.User{}
	fm := map[inmem.FollowKey]time.Time{}
	pr := inmem.NewPostRepository(pm)
	ur := inmem.NewUserRepository(um)
	fr := inmem.NewFollowRepository(fm, um)
	fns := repotest.TestFns{SetupTest: func() {
		for k := range pm {
			delete(pm, k)
		}
		for k := range um {
			delete(um, k)
		}
		for k := range fm {
			delete(fm, k)
		}
	}}
	suite.Run(t, repotest.NewPostSuite(pr, ur, fr, fns))
}
