package api

import (
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/db"

)

var dbActions db.Actions

func InitRoutes(p db.Actions) {
    router := gin.Default()
    dbActions = p
    router.GET("/recipes/:id", getRecipe)
    router.Run("localhost:8080")
}