package config

import (
    "fmt"
    "os"
    "context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func DBInit() {
    var err error
	DB, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}