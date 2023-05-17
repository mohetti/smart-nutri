package db

import (
     "github.com/mohetti/smart-nutri/api/models"
)



type Actions interface {
    OpenConnection() Actions
    CloseConnection()
    GetRecipe(id string, recipe *models.Recipe)
    GetFoods(id string, foods *[]models.Food)
}