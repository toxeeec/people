package inmem_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestInmemLikeSuite(t *testing.T) {
	lm := map[inmem.LikeKey]struct{}{}
	pm := map[uint]people.Post{}
	um := map[uint]people.User{}
	lr := inmem.NewLikeRepository(lm, pm, um)
	pr := inmem.NewPostRepository(pm)
	ur := inmem.NewUserRepository(um)
	fns := repotest.TestFns{SetupTest: func() {
		for k := range lm {
			delete(lm, k)
		}
		for k := range pm {
			delete(pm, k)
		}
	}}
	suite.Run(t, repotest.NewLikeSuite(lr, pr, ur, fns))
}
