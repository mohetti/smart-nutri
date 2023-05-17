package main

import (
    "github.com/mohetti/smart-nutri/api"
    "github.com/mohetti/smart-nutri/db"

)

func main() {
    postgres := db.Postgres.OpenConnection()
    api.RouterInit(postgres)
    defer postgres.CloseConnection()
}