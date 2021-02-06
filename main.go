package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	askAgain bool = true
	endHere  bool = false
)

type recipe struct {
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

var recipes []recipe

var replaceStr string

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	return strings.Replace(text, replaceStr, "", -1)
}

func setUp() {
	if runtime.GOOS == "windows" {
		replaceStr = "\r\n"
	} else {
		replaceStr = "\n"
	}
}

func main() {
	setUp()

	recipes = loadRecipes()
	for recipeMain() {

	}
}

func loadRecipes() []recipe {
	var recipes []recipe

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

func recipeMain() bool {
	fmt.Println(
		`===Options===
1) Generate random recipes
2) View recipes
3) Choose Recipes
4) Quit`,
	)

	text := getInput()
	choice, _ := strconv.Atoi(text)

	switch choice {
	case 1:
		generateRandomRecipes()
	case 2:
		viewRecipes()
	case 3:
		chooseRecipes()
	case 4:
		return endHere
	}
	return askAgain
}

func listRecipes() int {
	fmt.Println("Which recipe do you want to view?")
	for _, k := range recipes {
		fmt.Println(k.ID, "-", k.Name)
	}

	text := getInput()

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
		for _, recipe := range recipes {
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
	for _, k := range recipes {
		fmt.Println(k.ID, "-", k.Name)
	}
}

func chooseRecipes() {

}
