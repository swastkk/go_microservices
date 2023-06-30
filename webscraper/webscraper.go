package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Quote struct {
	Quote  string
	Author string
}

func main() {
	var quotes []Quote
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
		q := Quote{
			Quote:  quote,
			Author: author,
		}
		quotes = append(quotes, q)

	})

	c.Visit("https://quotes.toscrape.com/")

	csvFile, err := os.Create("quotes.csv")
	if err != nil {
		log.Fatalf("Failed in creating File: %s", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Header for the csv file!
	header := []string{"Quote", "Author"}
	err = csvWriter.Write(header)
	if err != nil {
		log.Fatal(err)
	}
	// Writing the data rows
	for _, quotes := range quotes {
		row := []string{quotes.Quote, quotes.Author}
		err = csvWriter.Write(row)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("CSV file created successfully")
}
