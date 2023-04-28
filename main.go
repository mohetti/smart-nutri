package main

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
	"github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/api/recipes"
)

func main() {
	var conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	router := gin.Default()
	recipes.Routing(router, conn)
    router.Run("localhost:8080")
}