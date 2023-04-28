package api

import (
    "fmt"
    "os"
    "context"
    "net/http"
    "github.com/jackc/pgx/v5"
    "github.com/gin-gonic/gin"
)

func handleRoute(dbConn *pgx.Conn) func(*gin.Context) {
    return func(c *gin.Context) {
        var name string
        var err = dbConn.QueryRow(context.Background(), "select name from recipes where id=$1", 1).Scan(&name)
        if err != nil {
            fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
            os.Exit(1)
        }

        c.IndentedJSON(http.StatusOK, name)
    }
}

func Connect(dbConn *pgx.Conn) {
    router := gin.Default()
    router.GET("/recipes", handleRoute(dbConn))

    router.Run("localhost:8080")
}