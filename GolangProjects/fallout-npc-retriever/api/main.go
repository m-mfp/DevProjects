package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

const URL = "https://fallout.fandom.com/wiki/Fallout_4_characters"

type npc struct {
	Name      string `json:"name"`
	Location  string `json:"location"`
	Doctor    bool   `json:"doctor"`
	Merchant  bool   `json:"merchant"`
	Companion bool   `json:"companion"`
	Essential bool   `json:"essential"`
}

func npcScrapper(g *gin.Context) {

	origin := g.GetHeader("Origin")
	fmt.Println("Request Origin:", origin)

	name := g.Param("name")

	var foundChar *npc
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error occurred while visiting %s: %v\n", r.Request.URL, err)
	})

	c.OnHTML(".mw-content-ltr.mw-parser-output", func(div *colly.HTMLElement) {
		div.ForEach("table", func(i int, table *colly.HTMLElement) {
			if i > 6 && i < 87 {
				table.ForEach("tbody tr", func(i int, tr *colly.HTMLElement) {
					if tr.ChildText("td:nth-child(1)") == name {
						fmt.Println("We found", name)
						foundChar = &npc{
							Name:      tr.ChildText("td:nth-child(1)"),
							Essential: tr.DOM.Find("td:nth-child(5) span").Length() == 3,
							Doctor:    tr.DOM.Find("td:nth-child(6) span").Length() == 3,
							Merchant:  tr.DOM.Find("td:nth-child(7) span").Length() == 3,
							Companion: tr.DOM.Find("td:nth-child(8) span").Length() == 3,
							Location:  tr.ChildText("td:nth-child(11)"),
						}

						g.IndentedJSON(http.StatusOK, foundChar)
						return
					}

				})
			}
		})
	})

	err := c.Visit(URL)
	if err != nil {
		log.Fatal(err)
		g.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to visit URL"})
		return
	}
}

func main() {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"}, // Allow your frontend URL
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	router.ForwardedByClientIP = true
	router.GET("/fallout-npc-scrapper/:name", npcScrapper)
	router.Run("localhost:12300")
}
