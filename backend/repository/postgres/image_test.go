package postgres_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/toxeeec/people/backend/repository/postgres"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestPostgresImageSuite(t *testing.T) {
	db, _ := sqlx.Connect("postgres", postgres.DSN)
	defer db.Close()
	ir := postgres.NewImageRepository(db)
	ur := postgres.NewUserRepository(db)
	pr := postgres.NewPostRepository(db)
	fns := repotest.TestFns{SetupTest: func() {
		db.MustExec("TRUNCATE image CASCADE")
	}}
	suite.Run(t, repotest.NewImageSuite(ir, ur, pr, fns))
}
