package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/m-mfp/skyrim-alchemy-scrapper/webscrapper"
)

func readCSV() (*csv.Reader, *os.File, error) {
	file, err := os.Open(webscrapper.CSVFileName)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file: %w", err)
	}

	reader := csv.NewReader(file)
	// defer file.Close()
	return reader, file, nil
}

func getEffects(ingredient string) ([]string, error) {

	reader, file, err := readCSV()

	if err != nil {
		return nil, fmt.Errorf("something went wrong with readCSV: %w", err)
	}
	defer file.Close()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading record: %w", err)
		}

		if strings.ToLower(record[0]) == ingredient {
			return record[1:], nil
		}
	}
	return nil, nil
}

func getIngredients(effect string) ([]string, error) {
	var ingredients []string
	reader, file, err := readCSV()
	defer file.Close()

	if err != nil {
		return nil, fmt.Errorf("something went wrong with readCSV: %w", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading record: %w", err)
		}

		for i := 1; i < 5; i++ {

			record[i] = strings.ToLower(record[i])
			if record[i] == strings.ToLower(effect) {
				ingredients = append(ingredients, record[0])
			}
		}
	}
	return ingredients, nil
}

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
