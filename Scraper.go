package main

import (
	"os"
	"time"

	"fmt"

	"encoding/json"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

func main() {
	allQuotes := make([]Quote, 0)

	c := colly.NewCollector()

	extensions.RandomUserAgent(c)

	extensions.Referer(c)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*quotes.toscrape.*",
		Parallelism: 2,
		Delay:       1 * time.Second,
		RandomDelay: 2 * time.Second,
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error: ", r.Request.URL, "Causa:", err)
	})

	c.OnHTML("div.quote", func(e *colly.HTMLElement) {
		fmt.Println(" ")
		Quote := Quote{
			Text:   e.ChildText("span.text"),
			Author: e.ChildText("small.author"),
		}
		allQuotes = append(allQuotes, Quote)

	})

	c.OnHTML("ul.pager", func(e *colly.HTMLElement) {
		nextPage := e.ChildAttr("li.next a", "href")
		if nextPage != "" {
			e.Request.Visit(e.Request.AbsoluteURL(nextPage))
		}

	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Status: %d, visit: %s\n", r.StatusCode, r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		file, _ := json.MarshalIndent(allQuotes, "", " ")
		_ = os.WriteFile("quotes.json", file, 0644)
		fmt.Println("Quotes saved to quotes.json")
	})

	c.Visit("https://quotes.toscrape.com/")

}

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}
