package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/types"
)

func getRecipe(c *gin.Context) {
   id := c.Param("id")
   recipe := types.Recipe{}

   routeConn.GetRecipe(id, &recipe)
   routeConn.GetFoods(id, &recipe.Foods)

   c.IndentedJSON(http.StatusOK, recipe)
}