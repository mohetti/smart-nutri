package api

import (
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/api/controllers"
)

func RouterInit() {
    router := gin.Default()
    routes(router)
    router.Run("localhost:8080")
}

func routes(router *gin.Engine) {
    router.GET("/recipes", controllers.Recipes)
}