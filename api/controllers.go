package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/api/models"
)

func getRecipe(c *gin.Context) {
   id := c.Param("id")
   recipe := models.Recipe{}

   dbActions.GetRecipe(id, &recipe)
   dbActions.GetFoods(id, &recipe.Foods)

   c.IndentedJSON(http.StatusOK, recipe)
}