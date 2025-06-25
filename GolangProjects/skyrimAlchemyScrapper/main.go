package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/m-mfp/skyrim-alchemy-scrapper/webscrapper"
)

func main() {

	if _, err := os.Stat(webscrapper.CSVFileName); os.IsNotExist(err) {

		err := webscrapper.DataCollection()
		if err != nil {
			log.Fatalf("Data collection failed: %v", err)
		}
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide an ingredient name or effect as a command-line argument.")
	}

	arg := os.Args[1]
	arg = strings.ToLower(arg)

	effects, err := getEffects(arg)
	if err != nil {
		log.Fatalf("Error with getEffects: %v", err)
	} else if effects != nil {
		fmt.Println("Effects:", effects)
		combinations, err := getCombinations(effects)
		if err != nil {
			log.Fatalf("Error with getCombinations: %v", err)
		}
		fmt.Println("Combinations:")
		for effect, ingredients := range combinations {
			fmt.Println("")
			fmt.Printf("-%s:\n", effect)
			for _, ingredient := range ingredients {
				fmt.Printf("-- %s\n", ingredient)
			}
		}
	}

	ingredients, err := getIngredients(arg)
	if err != nil {
		log.Fatalf("Error with getIngredients: %v", err)
	} else if ingredients != nil {
		fmt.Println("Ingredients:")
		for _, ingredient := range ingredients {
			fmt.Printf("- %s\n", ingredient)
		}
	}
}
