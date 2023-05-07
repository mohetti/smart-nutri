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

   var food = models.Food{}
// TODO: prevent sql injection on id
// TODO: format code
// TODO: implement repository pattern and clean up naming / file names and locations

   var foods, err2 = config.DB.Query(context.Background(), fmt.Sprintf(`select recipes_foods.id,
   %s
   from recipes_foods left join foods on recipes_foods.food_id = foods.id where recipes_foods.recipe_id=$1`, QueryString(food)), id)
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

   r.Foods = foodsSlice

   c.IndentedJSON(http.StatusOK, r)
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

func QueryString(u interface{}) string {
    val := reflect.ValueOf(u)
    s := ""

    for i := 0; i < val.NumField(); i++ {
    if (val.Type().Field(i).Name == "Id") {
        continue
    }
    s += val.Type().Field(i).Name

    if (i < val.NumField() - 1) {
    s += ", "
    }
    }
 return s
}