package main

import (
	"context"
	"fmt"
	"os"
	"net/http"
    "github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var dbConn *pgx.Conn

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// os.Getenv()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	dbConn = conn
	defer conn.Close(context.Background())

	getRecipes := func(c *gin.Context) {
	var name string
    	err = dbConn.QueryRow(context.Background(), "select name from recipes where id=$1", 1).Scan(&name)
    	if err != nil {
    		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    		os.Exit(1)
    	}

    	c.IndentedJSON(http.StatusOK, name)
	}



	router := gin.Default()
    router.GET("/recipes", getRecipes)

    router.Run("localhost:8080")
}