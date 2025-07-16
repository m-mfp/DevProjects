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

	var arg string

	if len(os.Args) < 2 {
		log.Fatal("Please provide an ingredient name or effect as a command-line argument.")
	} else {
		if os.Args[1] == "-p" || os.Args[1] == "-potion" {
			args := os.Args[2:]
			potionIngredients, err := createPotion(args)
			if err != nil {
				log.Fatalf("Error with createPotions: %v", err)
			}

			fmt.Println("Potion Ingredients are:")
			for _, ingredient := range potionIngredients {
				fmt.Println(ingredient)
			}
		} else {
			arg = strings.ToLower(os.Args[1])

			effects, err := getEffects(arg)
			if err != nil {
				log.Fatalf("Error with getEffects: %v", err)
			} else if effects != nil {
				fmt.Printf("Effects for %s:\n", arg)
				for _, effect := range effects {
					fmt.Printf("- %s\n", effect)
				}

				combinations, err := getCombinations(effects)
				if err != nil {
					log.Fatalf("Error with getCombinations: %v", err)
				}
				fmt.Println("\nCombinations:")
				for effect, ingredients := range combinations {
					fmt.Printf("\n-%s:\n", effect)
					for _, ingredient := range ingredients {
						fmt.Printf("-- %s\n", ingredient)
					}
				}
			}

			ingredients, err := getIngredients(arg)
			if err != nil {
				log.Fatalf("Error with getIngredients: %v", err)
			} else if ingredients != nil {
				fmt.Printf("Ingredients for %s:\n", arg)
				for _, ingredient := range ingredients {
					fmt.Printf("- %s\n", ingredient)
				}
			}
		}
	}
}
