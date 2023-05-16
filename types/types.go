package types

import (
    "github.com/mohetti/smart-nutri/api/models"
)

type Recipe struct {
    Id int
    Name string
    Foods []models.Food
}