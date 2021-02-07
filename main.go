package main

import (
	"github.com/go-recipes/common"
	"github.com/go-recipes/recipes"
)

func main() {
	common.SetUp()

	recipes.Recipes = recipes.LoadRecipes()
	for recipes.RecipeMain() {

	}
}
