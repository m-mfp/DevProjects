package main

import (
	"fmt"
	"log"
)

func getCombinations(effects []string) (map[string][]string, error) {
	combinations := make(map[string][]string)

	for _, effect := range effects {
		ingredients, err := getIngredients(effect)
		if err != nil {
			log.Fatalf("Error with getIngredients within getCombinations: %v", err)
		} else if ingredients == nil {
			return combinations, fmt.Errorf("empty ingredients, %v", err)
		}
		combinations[effect] = ingredients
	}
	return combinations, nil
}
