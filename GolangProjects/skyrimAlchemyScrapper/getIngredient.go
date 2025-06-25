package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/m-mfp/skyrim-alchemy-scrapper/webscrapper"
)

func getIngredients(effect string) ([]string, error) {
	var ingredients []string
	reader, file, err := webscrapper.ReadCSV()
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
