package controllers

import (
    "fmt"
    "os"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/config"
    "github.com/mohetti/smart-nutri/api/models"
)

var foodSelection = `select name, synonyms, category, density, reference_unit, kilojoule, kilocalories, fat_g,
saturated_fats_g, monounsaturated_fats_g, polyunsaturated_fats_g, cholesterol_mg, carbohydrates_g, sugar_g, starch_g,
dietary_fiber_g, protein_g, salt_nacl_g, alcohol_g, water_g, vitamin_a_activity_re_mug_re,
vitamin_a_activity_rae_mug_re, retinol_mug, betacarotin_activity_mug_bce, betacarotin_mug, vitamin_b1_mg,
vitamin_b2_mg, vitamin_b6_mg, vitamin_b12_mug, niacin_mg, folat_mug, vitamin_c_mg, vitamin_d_mug,
vitamin_e_activity_mg_ate, kalium_mg, natrium_mg, chlorid_mg, calcium_mg, magnesium_mg, phosphor_mg, iron_mg,
iodine_mug, zinc_mg, selenium_mug, foods.id from recipes_foods left join foods on recipes_foods.food_id = foods.id
where recipes_foods.recipe_id=$1`

type recipe struct {
    Id int
    Name string
    Foods []models.Food
}

func GetRecipe(c *gin.Context) {
   id := c.Param("id")
   recipe := recipe{}

   if recipeErr := config.DB.Get(&recipe, "select id, name from recipes where id=$1", id); recipeErr != nil {
       fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", recipeErr)
       os.Exit(1)
   }

   if foodsErr := config.DB.Select(&recipe.Foods, foodSelection, id); foodsErr != nil {
          fmt.Fprintf(os.Stderr, "Query failed: %v\n", foodsErr)
          os.Exit(1)
      }

   c.IndentedJSON(http.StatusOK, recipe)
}