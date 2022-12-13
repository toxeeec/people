package postgres_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/toxeeec/people/backend/repository/postgres"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestPostgresFollowSuite(t *testing.T) {
	db, _ := sqlx.Connect("postgres", postgres.DSN)
	defer db.Close()
	fr := postgres.NewFollowRepository(db)
	ur := postgres.NewUserRepository(db)
	fns := repotest.TestFns{SetupTest: func() {
		db.MustExec("TRUNCATE user_profile CASCADE")
	}}
	suite.Run(t, repotest.NewFollowSuite(fr, ur, fns))
}
