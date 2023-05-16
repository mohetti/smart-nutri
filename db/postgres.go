package db

import (
    "fmt"
    "os"
    _ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/mohetti/smart-nutri/api/models"
    "github.com/mohetti/smart-nutri/types"
)

var foodSelection = `select name, synonyms, category, density, reference_unit, kilojoule, kilocalories, fat_g,
saturated_fats_g, monounsaturated_fats_g, polyunsaturated_fats_g, cholesterol_mg, carbohydrates_g, sugar_g, starch_g,
dietary_fiber_g, protein_g, salt_nacl_g, alcohol_g, water_g, vitamin_a_activity_re_mug_re,
vitamin_a_activity_rae_mug_re, retinol_mug, betacarotin_activity_mug_bce, betacarotin_mug, vitamin_b1_mg,
vitamin_b2_mg, vitamin_b6_mg, vitamin_b12_mug, niacin_mg, folat_mug, vitamin_c_mg, vitamin_d_mug,
vitamin_e_activity_mg_ate, kalium_mg, natrium_mg, chlorid_mg, calcium_mg, magnesium_mg, phosphor_mg, iron_mg,
iodine_mug, zinc_mg, selenium_mug, foods.id from recipes_foods left join foods on recipes_foods.food_id = foods.id
where recipes_foods.recipe_id=$1`

type Postgres struct {
    DB *sqlx.DB
}

var PostgresDb = Postgres{}

func (p Postgres) GetRecipe(id string, recipe *types.Recipe) {
    if recipeErr := p.DB.Get(recipe, "select id, name from recipes where id=$1", id); recipeErr != nil {
       fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", recipeErr)
       os.Exit(1)
   }
}

func (p Postgres) GetFoods(id string, foods *[]models.Food) {
    if foodsErr := p.DB.Select(foods, foodSelection, id); foodsErr != nil {
              fmt.Fprintf(os.Stderr, "Query failed: %v\n", foodsErr)
              os.Exit(1)
          }
}

func DBInit() *Postgres {
    var err error
	PostgresDb.DB, err = sqlx.Connect("postgres", os.Getenv("POSTGRES_CREDENTIALS"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &PostgresDb
}