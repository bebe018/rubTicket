package concertInfo

import (
	"github.com/gocolly/colly"
)

var url string = GetUrl()

func GetConcertBookPage() string {
	var concertBookMap = make(map[string]string)

	settings := GetJSON()

	c := colly.NewCollector()
	c.OnHTML(".gridc", func(h *colly.HTMLElement) {
		// concertDay :=
		link := h.ChildAttr(".btn", "data-href")
		date := h.ChildText("td")[0:10]
		concertBookMap[date] = link
	})
	c.Visit(url)

	return concertBookMap[settings.Date]

}
