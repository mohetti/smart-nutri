package db

import (
    "github.com/mohetti/smart-nutri/types"
)



type dbActions interface {
    GetRecipe(id string, recipe *types.Recipe)
    GetFoods(id string, foods *types.Recipe)
}