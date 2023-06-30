package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Repo struct {
	Name        string
	Desc        string
	Technology  string
	LastUpdated string
}

func main() {
	var quotes []Repo
	c := colly.NewCollector()
	colly.AllowedDomains("github.com")
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/114.0")
		fmt.Println("URL to Scrape data: ", r.URL)
	})

	// If any error occurs
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err.Error())
	})

	//
	c.OnHTML(".d-inline-block", func(h *colly.HTMLElement) {
		div := h.DOM
		name := div.Find("[itemprop='name codeRepository']").First().Text()
		desc := div.Find("[itemprop='description']").First().Text()
		lang := div.Find("[itemprop='programmingLanguage']").First().Text()
		time := div.Find("relative-time").First().Text()
		q := Repo{
			Name:        name,
			Desc:        desc,
			Technology:  lang,
			LastUpdated: time,
		}
		quotes = append(quotes, q)
	})

	c.Visit("https://github.com/swastkk?tab=repositories")

	// Creating a CSV file to store the Scraped Data
	csvFile, err := os.Create("data.csv")
	if err != nil {
		log.Fatalf("Failed in creating File: %s", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Header for the csv file!
	header := []string{"Name", "Description", "Technology", "Last Updated"}
	err = csvWriter.Write(header)
	if err != nil {
		log.Fatal(err)
	}
	// Writing the data rows
	for _, q := range quotes {
		row := []string{q.Name, q.Desc, q.Technology, q.LastUpdated}
		err = csvWriter.Write(row)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("CSV file created successfully")
}
