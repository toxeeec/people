package inmem_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestInmemImageSuite(t *testing.T) {
	im := map[uint]people.Image{}
	um := map[uint]people.User{}
	ir := inmem.NewImageRepository(im)
	ur := inmem.NewUserRepository(um)
	fns := repotest.TestFns{SetupTest: func() {
		for k := range im {
			delete(im, k)
		}
		for k := range um {
			delete(um, k)
		}
	}}
	suite.Run(t, repotest.NewImageSuite(ir, ur, fns))
}
