package main

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
	"github.com/mohetti/smart-nutri/api"
)

func main() {
	var conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	api.Connect(conn)
}