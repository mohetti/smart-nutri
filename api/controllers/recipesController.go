package controllers

import (
    "fmt"
    "os"
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/mohetti/smart-nutri/config"
    "github.com/mohetti/smart-nutri/api/models"
)

type recipe struct {
    Id int
    Name string
    Foods []models.Food
}

func Recipes(c *gin.Context) {
   id := c.Param("id")
   r := recipe{}
   var err = config.DB.QueryRow(context.Background(), "select id, name from recipes where id=$1", id).Scan(&r.Id, &r.Name)
   if err != nil {
       fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
       os.Exit(1)
   }

   var foods, err2 = config.DB.Query(context.Background(), `select
   recipes_foods.id, food_id, recipe_id, name, synonyms, category, density, reference_unit, kilojoule, kilocalories,
   fat_g, saturated_fats_g, monounsaturated_fats_g, polyunsaturated_fats_g, cholesterol_mg, carbohydrates_g,
   sugar_g, starch_g, dietary_fiber_g, protein_g, salt_nacl_g, alcohol_g, water_g, vitamin_a_activity_re_mug_re,
   vitamin_a_activity_rae_mug_re, retinol_mug, betacarotin_activity_mug_bce, betacarotin_mug, vitamin_b1_mg,
   vitamin_b2_mg, vitamin_b6_mg, vitamin_b12_mug, niacin_mg, folat_mug, vitamin_c_mg, vitamin_d_mug, vitamin_e_activity_mg_ate,
   kalium_mg, natrium_mg, chlorid_mg, calcium_mg, magnesium_mg, phosphor_mg, iron_mg, iodine_mug, zinc_mg, selenium_mug
   from recipes_foods left join foods on recipes_foods.food_id = foods.id where recipes_foods.recipe_id=$1`, id)
   if err2 != nil {
          fmt.Fprintf(os.Stderr, "Query failed: %v\n", err2)
          os.Exit(1)
      }

   defer foods.Close()

   var foodsSlice []models.Food

   for foods.Next() {
   var food models.Food
   foods.Scan(
                 &food.Id,
                 &food.Food_id,
                 &food.Recipe_id,
                 &food.Name,
                 &food.Synonyms,
                 &food.Category,
                 &food.Density,
                 &food.Reference_unit,
                 &food.Kilojoule,
                 &food.Kilocalories,
                 &food.Fat_g,
                 &food.Saturated_fats_g,
                 &food.Monounsaturated_fats_g,
                 &food.Polyunsaturated_fats_g,
                 &food.Cholesterol_mg,
                       &food.Carbohydrates_g,
                       &food.Sugar_g,
                       &food.Starch_g,
                       &food.Dietary_fiber_g,
                       &food.Protein_g,
                       &food.Salt_nacl_g,
                       &food.Alcohol_g,
                       &food.Water_g,
                       &food.Vitamin_a_activity_re_mug_re,
                       &food.Vitamin_a_activity_rae_mug_re,
                       &food.Retinol_mug,
                       &food.Betacarotin_activity_mug_bce,
                       &food.Betacarotin_mug,
                       &food.Vitamin_b1_mg,
                       &food.Vitamin_b2_mg,
                       &food.Vitamin_b6_mg,
                       &food.Vitamin_b12_mug,
                       &food.Niacin_mg,
                       &food.Folat_mug,
                       &food.Vitamin_c_mg,
                       &food.Vitamin_d_mug,
                       &food.Vitamin_e_activity_mg_ate,
                       &food.Kalium_mg,
                       &food.Natrium_mg,
                       &food.Chlorid_mg,
                       &food.Calcium_mg,
                       &food.Magnesium_mg,
                       &food.Phosphor_mg,
                       &food.Iron_mg,
                       &food.Iodine_mug,
                       &food.Zinc_mg,
                       &food.Selenium_mug)
   foodsSlice = append(foodsSlice, food)
   }

   r.Foods = foodsSlice

   c.IndentedJSON(http.StatusOK, r)
}