package search

import (
	"github.com/gocolly/colly"
)

var concertBookPage = make(map[string]string)

func GetConcertBookPage(url string) map[string]string {
	c := colly.NewCollector()
	c.OnHTML(".gridc", func(h *colly.HTMLElement) {
		// concertDay :=
		link := h.ChildAttr(".btn", "data-href")
		date := h.ChildText("td")[0:10]
		concertBookPage[date] = link
	})
	c.Visit(url)
	return concertBookPage
}
