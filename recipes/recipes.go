package recipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/go-recipes/common"
)

// Recipe - Defines a structure to unmarshall a json into
type Recipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Ingredients []struct {
		Name   string `json:"name"`
		Amount int    `json:"amount"`
		Units  string `json:"units"`
	} `json:"ingredients"`
	Steps []struct {
		Step   int    `json:"step"`
		Title  string `json:"title"`
		Text   string `json:"text"`
		Timers []struct {
			Title string `json:"title"`
			Time  int    `json:"time"`
		} `json:"timers"`
	} `json:"steps"`
}

// Recipes - The stored list of recipes
var Recipes []Recipe

// LoadRecipes - Loads in and unmarshalls the json
func LoadRecipes() []Recipe {
	var recipes []Recipe

	jsonFile, err := os.Open("recipes.json")

	if err != nil {
		fmt.Println(err)
		return recipes
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Println("Successfully Opened recipes.json")
	json.Unmarshal(byteValue, &recipes)

	return recipes
}

func listRecipes() int {
	fmt.Println("Which recipe do you want to view?")
	for _, k := range Recipes {
		fmt.Println(k.ID, "-", k.Name)
	}

	text := common.GetInput()

	choice, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(choice)
	return choice
}

func getShoppingList(ids ...int) {
	var ingredients = make(map[string]int)
	for _, id := range ids {
		for _, recipe := range Recipes {
			if recipe.ID == id {
				for _, ingredient := range recipe.Ingredients {
					ingredients[ingredient.Name] += ingredient.Amount
				}
			}
		}
	}

	fmt.Println(ingredients)
}

func generateRandomRecipes() {

}

func viewRecipes() {
	for _, k := range Recipes {
		fmt.Println(k.ID, "-", k.Name)
	}
}

func chooseRecipes() {

}

// RecipeMain - The main function that determines how the user can proceed
func RecipeMain() bool {
	fmt.Println(
		`===Options===
1) Generate random recipes
2) View recipes
3) Choose Recipes
4) Quit`,
	)

	text := common.GetInput()
	choice, _ := strconv.Atoi(text)

	switch choice {
	case 1:
		generateRandomRecipes()
	case 2:
		viewRecipes()
	case 3:
		chooseRecipes()
	case 4:
		return common.EndHere
	}
	return common.AskAgain
}
