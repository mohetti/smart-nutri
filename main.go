package main

import (
    "github.com/mohetti/smart-nutri/api"
    "github.com/mohetti/smart-nutri/db"

)

func main() {
    postgres := db.DBInit()
    api.RouterInit(postgres)
    defer postgres.DB.Close()
}