// Package scraper implements the build page scraper as a single function.
package scraper

import (
	"context"
	"fmt"

	"github.com/gocolly/colly"
)

// Scraper defines a type that other Scrapers implement
type Scraper func(ctx context.Context) ([]interface{}, error)

// Scrape scrapes for starcraft builds from
func Scrape() {
	c := colly.NewCollector()

	// main content area
	c.OnHTML(".col-md-9", func(h *colly.HTMLElement) {
		fmt.Printf("build list: %+v\n", h)
		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			fmt.Printf("\n --- build: %+v\n", h)
		})
	})

	c.Visit("https://lotv.spawningtool.com/build/")
}
