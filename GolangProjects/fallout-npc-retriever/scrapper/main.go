package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"golang.org/x/time/rate"
)

const URL = "https://fallout.fandom.com/"

type npc struct {
	Name      string `json:"name"`
	Location  string `json:"location"`
	Doctor    bool   `json:"doctor"`
	Merchant  bool   `json:"merchant"`
	Companion bool   `json:"companion"`
	Essential bool   `json:"essential"`
	Photo     string `json:"photo"`
	Brief     string `json:"brief"`
}

func npcScrapper(g *gin.Context) {

	origin := g.GetHeader("Origin")
	fmt.Println("Request Origin:", origin)

	name := g.Param("name")

	if err := validateName(name); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundChar *npc
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println("Got a response from", r.Request.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error occurred while visiting %s: %v\n", r.Request.URL.String(), err)
		g.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
	})

	c.OnHTML(".mw-content-ltr.mw-parser-output", func(div *colly.HTMLElement) {
		var npcURL string
		div.ForEach("table", func(i int, table *colly.HTMLElement) {
			if i > 6 && i < 87 {
				table.ForEach("tbody tr", func(i int, tr *colly.HTMLElement) {
					if tr.ChildText("td:nth-child(1)") == name {
						npcURL = tr.ChildAttr("td:nth-child(1) a", "href")

						foundChar = &npc{
							Name:      tr.ChildText("td:nth-child(1)"),
							Essential: tr.DOM.Find("td:nth-child(5) span").Length() == 3,
							Doctor:    tr.DOM.Find("td:nth-child(6) span").Length() == 3,
							Merchant:  tr.DOM.Find("td:nth-child(7) span").Length() == 3,
							Companion: tr.DOM.Find("td:nth-child(8) span").Length() == 3,
							Location:  tr.ChildText("td:nth-child(11)"),
						}
					}

				})
			}
		})

		// npc url
		if npcURL != "" {
			npcCollector := colly.NewCollector()
			npcCollector.OnHTML(".mw-content-ltr.mw-parser-output", func(el *colly.HTMLElement) {
				npcPhoto := el.ChildAttr("figure img", "src")

				if npcPhoto != "" {
					foundChar.Photo = npcPhoto
				}

				var npcBrief string
				el.DOM.Find("div#toc").Each(func(i int, tocEl *goquery.Selection) {
					if tocEl.Prev().Length() > 0 && tocEl.Prev().Get(0).Data == "p" {
						npcBrief = tocEl.Prev().Text()
					}
				})

				if npcBrief != "" {
					foundChar.Brief = npcBrief
				}

			})

			npcCollector.Visit(URL + npcURL)
		} else {
			g.IndentedJSON(http.StatusNotFound, gin.H{"error": "NPC not found"})
		}
		g.JSON(http.StatusOK, foundChar)
	})

	err := c.Visit(URL + "wiki/Fallout_4_characters")
	if err != nil {
		log.Fatal(err)
		g.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to visit URL"})
		return
	}
}

func validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if len(name) > 25 {
		return fmt.Errorf("name cannot exceed 25 characters")
	}

	validNamePattern := `^[a-zA-Z\s'-]+$`
	matched, err := regexp.MatchString(validNamePattern, name)
	if err != nil || !matched {
		return fmt.Errorf("name contains invalid characters")
	}

	return nil
}

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 4)
	return func(c *gin.Context) {

		if limiter.Allow() {
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Limite exceed",
			})
		}

	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"}, // Allow your frontend URL
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	router.Use(RateLimiter())

	router.ForwardedByClientIP = true
	router.GET("/fallout-npc-scrapper/:name", npcScrapper)
	router.Run("localhost:12300")
}
