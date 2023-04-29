package controllers

import (
    "fmt"
    "os"
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/config"
)

func Recipes(c *gin.Context) {
   var name string
   var err = config.DB.QueryRow(context.Background(), "select name from recipes where id=$1", 1).Scan(&name)
   if err != nil {
       fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
       os.Exit(1)
   }

   c.IndentedJSON(http.StatusOK, name)
}