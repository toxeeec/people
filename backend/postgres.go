package people

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	host     = "host.docker.internal"
	port     = os.Getenv("POSTGRES_PORT")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
	dsn      = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

func PostgresConnect() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dsn)
}
