package postgres_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/toxeeec/people/backend/repository/postgres"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestPostgresLikeSuite(t *testing.T) {
	db, _ := sqlx.Connect("postgres", postgres.DSN)
	defer db.Close()
	lr := postgres.NewLikeRepository(db)
	pr := postgres.NewPostRepository(db)
	ur := postgres.NewUserRepository(db)
	fns := repotest.TestFns{SetupTest: func() {
		db.MustExec("TRUNCATE user_profile CASCADE")
		db.MustExec("TRUNCATE post CASCADE")
		db.MustExec("TRUNCATE post_like CASCADE")
	}}
	suite.Run(t, repotest.NewLikeSuite(lr, pr, ur, fns))
}
