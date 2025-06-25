package main

import (
	"fmt"
	"slices"
	"strings"
)

func createPotion(args []string) ([]string, error) {
	allIngredients := make(map[string][]string)

	for _, arg := range args {
		ingredients, err := getIngredients(strings.ToLower(arg))
		if err != nil {
			return nil, fmt.Errorf("something went wrong with getIngredients: %w", err)
		} else if ingredients != nil {
			allIngredients[arg] = ingredients
		}
	}

	var potionIngredients []string

	for i := 1; i < len(args); i++ {
		for _, ingredient := range allIngredients[args[0]] {
			if slices.Contains(allIngredients[args[i]], ingredient) {
				potionIngredients = append(potionIngredients, ingredient)
			}
		}
	}

	return potionIngredients, nil
}
