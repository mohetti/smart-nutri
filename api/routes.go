package api

import (
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/api/controllers"
)

func RouterInit() {
    router := gin.Default()
    router.GET("/recipes/:id", controllers.GetRecipe)
    router.Run("localhost:8080")
}