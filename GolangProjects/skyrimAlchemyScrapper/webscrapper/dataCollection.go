package webscrapper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func DataCollection(c *colly.Collector, url string) ([]Ingredient, error) {
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
		return nil, fmt.Errorf("error visiting the URL: %w", err)
	}

	return ingredients, nil
}
