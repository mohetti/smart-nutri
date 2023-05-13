package main

import (
    "github.com/mohetti/smart-nutri/config"
    "github.com/mohetti/smart-nutri/api"
)

func main() {
    config.DBInit()
    api.RouterInit()
    defer config.DB.Close()
}