package postgres_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/toxeeec/people/backend/repository/postgres"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestPostgresTokenSuite(t *testing.T) {
	db, _ := sqlx.Connect("postgres", postgres.DSN)
	defer db.Close()
	tr := postgres.NewTokenRepository(db)
	ur := postgres.NewUserRepository(db)
	fns := repotest.TestFns{SetupTest: func() {
		db.MustExec("TRUNCATE user_profile CASCADE")
	}}
	suite.Run(t, repotest.NewTokenSuite(tr, ur, fns))
}
