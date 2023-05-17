package api

import (
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/db"

)

var routeConn db.Actions

func RouterInit(dbActions db.Actions) {
    router := gin.Default()
    routeConn = dbActions
    router.GET("/recipes/:id", getRecipe)
    router.Run("localhost:8080")
}