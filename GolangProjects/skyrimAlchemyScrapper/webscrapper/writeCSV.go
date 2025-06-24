package webscrapper

import (
	"encoding/csv"
	"fmt"
	"os"
)

const csvFileName = "ingredients.csv"

type Ingredient struct {
	Title   string
	Effects []string
}

var csvHeaders = []string{"Ingredient", "Effect One", "Effect Two", "Effect Three", "Effect Four"}

func WriteCSV(ingredients []Ingredient) error {
	file, err := os.Create(csvFileName)
	if err != nil {
		return fmt.Errorf("error creating CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(csvHeaders); err != nil {
		return fmt.Errorf("error writing header to CSV: %w", err)
	}

	for _, ingredient := range ingredients {
		record := []string{ingredient.Title}
		record = append(record, ingredient.Effects...)

		for len(record) < len(csvHeaders) {
			record = append(record, "")
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing record to CSV: %w", err)
		}
	}
	return nil
}
