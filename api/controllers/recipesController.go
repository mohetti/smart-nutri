package controllers

import (
    "fmt"
    "os"
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/config"
)

type recipe struct {
    Id int
    Name string
}

func Recipes(c *gin.Context) {
   r := recipe{}
   var err = config.DB.QueryRow(context.Background(), "select id, name from recipes where id=$1", 1).Scan(&r.Id, &r.Name)
   if err != nil {
       fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
       os.Exit(1)
   }

   c.IndentedJSON(http.StatusOK, r)
}