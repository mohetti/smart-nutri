package db

import (
    "github.com/mohetti/smart-nutri/types"
     "github.com/mohetti/smart-nutri/api/models"
)



type Actions interface {
    OpenConnection() Actions
    CloseConnection()
    GetRecipe(id string, recipe *types.Recipe)
    GetFoods(id string, foods *[]models.Food)
}