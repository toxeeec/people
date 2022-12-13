package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/toxeeec/people/backend/http"
	"github.com/toxeeec/people/backend/repository/postgres"
)

func main() {
	db, err := sqlx.Connect("postgres", postgres.DSN)
	if err != nil {
		log.Fatal(err)
	}

	v := validator.New()
	e := http.NewServer(db, v)
	e.Logger.Fatal(e.Start(":8000"))
}
