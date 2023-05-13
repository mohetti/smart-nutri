package config

import (
    "fmt"
    "os"
    _ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func DBInit() {
    var err error
	DB, err = sqlx.Connect("postgres", os.Getenv("POSTGRES_CREDENTIALS"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}