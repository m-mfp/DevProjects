package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/m-mfp/skyrim-alchemy-scrapper/webscrapper"
)

const csvFileName = "ingredients.csv"

func readCSV() (*csv.Reader, error) {
	file, err := os.Open(csvFileName)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error reading header: %w", err)
	}

	return reader, nil
}

func getEffects(arg string) ([]string, error) {

	reader, err := readCSV()
	if err == nil {
		for {
			record, err := reader.Read()
			if err != nil {
				return nil, fmt.Errorf("error reading record: %w", err)
			} else if err == io.EOF {
				break
			}

			if strings.ToLower(record[0]) == arg {
				effects := record[1:]
				return effects, nil
			}
		}
	}

	return nil, fmt.Errorf("something went wrong with geteffects: %w", err)
}

func getIngredients(arg string) ([]string, error) {

	var ingredients []string
	reader, err := readCSV()
	if err != nil {
		return nil, fmt.Errorf("something went wrong with getIngredients: %w", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading record: %w", err)
		}

		for _, effect := range record[1:] {
			if strings.ToLower(effect) == arg {
				ingredients = append(ingredients, record[0])

			}
			return ingredients, nil
		}
	}
	return nil, nil
}

func main() {
	url := "https://elderscrolls.fandom.com/wiki/Ingredients_(Skyrim)"
	collector := colly.NewCollector()

	if _, err := os.Stat(csvFileName); os.IsNotExist(err) {

		ingredients, err := webscrapper.DataCollection(collector, url)
		if err != nil {
			log.Fatalf("Data collection failed: %v", err)
		}

		if err := webscrapper.WriteCSV(ingredients); err != nil {
			log.Fatalf("Writing CSV failed: %v", err)
		}
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide an ingredient name or effect as a command-line argument.")
	}
	arg := os.Args[1]
	arg = strings.ToLower(arg)

	effects, _ := getEffects(arg)
	fmt.Println("Effects:", effects)

	ingredientsList, _ := getIngredients(arg)
	fmt.Println("Ingredients:", ingredientsList)
}
