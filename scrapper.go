package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Struct to hold the scraped car information
type Car struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
	Image string `json:"image"`
}

// Function to fetch the page
func fetchPage(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching the URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: status code %d when fetching %s", resp.StatusCode, url)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading document: %v", err)
	}
	return doc, nil
}

// Function to scrape the pages
func scrapeAllPages() {
	baseUrl := "https://999.md/ro/list/transport/cars?page="
	page := 1
	noAnnouncements := false

	// Get current date and format filename
	today := time.Now()
	formattedDate := fmt.Sprintf("%d_%02d_%02d", today.Year(), today.Month(), today.Day())
	resultsFile := fmt.Sprintf("scrapped_%s.json", formattedDate)

	// Initialize the JSON file
	file, err := os.Create(resultsFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Write the opening bracket for the JSON array
	file.WriteString("[")

	firstPage := true

	for !noAnnouncements {
		url := fmt.Sprintf("%s%d", baseUrl, page)
		doc, err := fetchPage(url)
		if err != nil {
			fmt.Printf("Failed to load page %d: %v\n", page, err)
			break
		}

		// Find the car listings
		announcements := doc.Find(".ads-list-photo.large-photo .ads-list-photo-item")
		if announcements.Length() == 0 {
			noAnnouncements = true
			fmt.Printf("No announcements found on page %d. Stopping.\n", page)
			break
		}

		var pageResults []Car

		announcements.Each(func(i int, s *goquery.Selection) {
			classes, _ := s.Attr("class")
			if strings.Contains(classes, "js-booster-inline") || strings.Contains(classes, "is-adsense") {
				return
			}

			titleElement := s.Find(".ads-list-photo-item-title a")
			priceElement := s.Find(".ads-list-photo-item-price-wrapper")

			title := strings.TrimSpace(titleElement.Text())
			price := strings.TrimSpace(priceElement.Text())
			href, exists := titleElement.Attr("href")
			id := ""
			if exists {
				parts := strings.Split(href, "/")
				id = parts[len(parts)-1]
			}
			imgElement := s.Find(".ads-list-photo-item-thumb img")
			imgUrl, _ := imgElement.Attr("src")

			if title == "" || id == "" {
				return
			}

			pageResults = append(pageResults, Car{
				ID:    id,
				Title: title,
				Price: strings.ReplaceAll(price, "\u00a0", " "),
				Image: imgUrl,
			})
		})

		if len(pageResults) > 0 {
			// Write page results to JSON file
			data, err := json.MarshalIndent(pageResults, "", "  ")
			if err != nil {
				fmt.Printf("Error marshaling JSON: %v\n", err)
				return
			}

			if !firstPage {
				file.WriteString(",\n")
			}

			file.Write(data)
			firstPage = false
		}

		fmt.Printf("Page %d processed.\n", page)
		page++
	}

	// Finalize the JSON file by closing the array
	file.WriteString("]")
	fmt.Printf("Results have been saved to %s\n", resultsFile)
}

func main() {
	scrapeAllPages()
}
