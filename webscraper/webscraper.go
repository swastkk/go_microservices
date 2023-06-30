package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	colly.AllowedDomains("quotes.toscrape.com")
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/114.0")
		fmt.Println("Visting a URL", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode) // To check the status code of the response of the website which is to be scraped!
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err.Error())
	})
	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		fmt.Printf("Quote: %v\n", quote)
		fmt.Printf("Author: %v\n", author)

	})
	// c.OnHTML(".text", func(h *colly.HTMLElement) {
	// 	fmt.Println("Quote: ", h.Text)
	// })
	// c.OnHTML(".author", func(h *colly.HTMLElement) {
	// 	fmt.Println("Author: ", h.Text)
	// })
	c.Visit("https://quotes.toscrape.com/")
}
