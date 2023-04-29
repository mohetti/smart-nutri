package main

import (
    "context"
    "github.com/mohetti/smart-nutri/config"
    "github.com/mohetti/smart-nutri/api"
)

func main() {
    config.DBInit()
    api.RouterInit()
    defer config.DB.Close(context.Background())

}