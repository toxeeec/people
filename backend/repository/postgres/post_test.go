package postgres_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/toxeeec/people/backend/repository/postgres"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestPostgresPostSuite(t *testing.T) {
	db, _ := sqlx.Connect("postgres", postgres.DSN)
	defer db.Close()
	pr := postgres.NewPostRepository(db)
	ur := postgres.NewUserRepository(db)
	fr := postgres.NewFollowRepository(db)
	fns := repotest.TestFns{SetupTest: func() {
		db.MustExec("TRUNCATE post CASCADE")
	}}
	suite.Run(t, repotest.NewPostSuite(pr, ur, fr, fns))
}
