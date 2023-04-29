package config

import (
    "fmt"
     "os"
     "context"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func DBInit() {
    var err error
	DB, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}