package controllers

import (
    "fmt"
    "os"
    "context"
    "net/http"
    "reflect"
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/config"
    "github.com/mohetti/smart-nutri/api/models"
)

var foodSelection = `select name, synonyms, category, density, reference_unit, kilojoule, kilocalories, fat_g,
saturated_fats_g, monounsaturated_fats_g, polyunsaturated_fats_g, cholesterol_mg, carbohydrates_g, sugar_g, starch_g,
dietary_fiber_g, protein_g, salt_nacl_g, alcohol_g, water_g, vitamin_a_activity_re_mug_re,
vitamin_a_activity_rae_mug_re, retinol_mug, betacarotin_activity_mug_bce, betacarotin_mug, vitamin_b1_mg,
vitamin_b2_mg, vitamin_b6_mg, vitamin_b12_mug, niacin_mg, folat_mug, vitamin_c_mg, vitamin_d_mug,
vitamin_e_activity_mg_ate, kalium_mg, natrium_mg, chlorid_mg, calcium_mg, magnesium_mg, phosphor_mg, iron_mg float32,
iodine_mug, zinc_mg, selenium_mug, foods.id from recipes_foods left join foods on recipes_foods.food_id = foods.id
where recipes_foods.recipe_id=$1`

func GetRecipe(c *gin.Context) {
   id := c.Param("id")
   recipe := models.Recipe{}
   var err = config.DB.QueryRow(context.Background(), "select id, name from recipes where id=$1", id).Scan(&recipe.Id, &recipe.Name)

   if err != nil {
       fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
       os.Exit(1)
   }

// TODO: prevent sql injection on id
// TODO: format code
// TODO: implement repository pattern and clean up naming / file names and locations

   var foods, err2 = config.DB.Query(context.Background(), foodSelection, id)
   if err2 != nil {
          fmt.Fprintf(os.Stderr, "Query failed: %v\n", err2)
          os.Exit(1)
      }

   defer foods.Close()

   var foodsSlice []models.Food

   for foods.Next() {
   var food models.Food
   t := StructForScan(&food)
   foods.Scan(t...)
   foodsSlice = append(foodsSlice, food)
   }

   recipe.Foods = foodsSlice

   c.IndentedJSON(http.StatusOK, recipe)
}

func StructForScan(u interface{}) []interface{} {
    val := reflect.ValueOf(u).Elem()
    v := make([]interface{}, val.NumField())

    for i := 0; i < val.NumField(); i++ {
        valueField := val.Field(i)
        v[i] = valueField.Addr().Interface()
    }
    return v
}