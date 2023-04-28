package recipes

import (
    "fmt"
    "os"
    "context"
    "net/http"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Routing(router *gin.Engine, dbConn *pgx.Conn) {
    router.GET("/recipes", handleRoute(dbConn))

}

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

