package webscrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

const CSVFileName = "ingredients.csv"

type Ingredient struct {
	Title   string
	Effects []string
}

var csvHeaders = []string{"Ingredient", "Effect One", "Effect Two", "Effect Three", "Effect Four"}

func DataCollection() error {
	url := "https://elderscrolls.fandom.com/wiki/Ingredients_(Skyrim)"
	c := colly.NewCollector()

	var ingredients []Ingredient

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error occurred while visiting %s: %v\n", r.Request.URL, err)
	})

	c.OnHTML(".wikitable.sortable.highlight tbody tr", func(row *colly.HTMLElement) {
		ingredient := Ingredient{
			Title: row.ChildText("th a"),
		}
		for i := 2; i < 6; i++ {
			effect := row.ChildText(fmt.Sprintf("td:nth-child(%d) a", i))
			if effect != "" {
				ingredient.Effects = append(ingredient.Effects, effect)
			}
		}

		if ingredient.Title != "" {
			ingredients = append(ingredients, ingredient)
		}
	})

	if err := c.Visit(url); err != nil {
		return fmt.Errorf("error visiting the URL: %w", err)
	}

	err := WriteCSV(ingredients)
	if err != nil {
		return fmt.Errorf("error writing record to CSV: %w", err)
	}
	return nil
}

func WriteCSV(ingredients []Ingredient) error {
	file, err := os.Create(CSVFileName)
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

func ReadCSV() (*csv.Reader, *os.File, error) {
	file, err := os.Open(CSVFileName)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file: %w", err)
	}

	reader := csv.NewReader(file)
	return reader, file, nil
}
