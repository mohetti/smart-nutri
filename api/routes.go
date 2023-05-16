package api

import (
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/db"

)

var RouteConn  *db.Postgres

func RouterInit(postgres *db.Postgres) {
    router := gin.Default()
    RouteConn = postgres
    router.GET("/recipes/:id", GetRecipe)
    router.Run("localhost:8080")
}