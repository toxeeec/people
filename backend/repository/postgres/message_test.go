package postgres_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/toxeeec/people/backend/repository/postgres"
	"github.com/toxeeec/people/backend/repository/repotest"
)

func TestPostgresMessageSuite(t *testing.T) {
	db, _ := sqlx.Connect("postgres", postgres.DSN)
	defer db.Close()
	mr := postgres.NewMessageRepository(db)
	ur := postgres.NewUserRepository(db)
	fns := repotest.TestFns{SetupTest: func() {
		db.MustExec("TRUNCATE message CASCADE")
	}}
	suite.Run(t, repotest.NewMessageSuite(mr, ur, fns))
}
