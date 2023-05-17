package concertInfo

import (
	"github.com/gocolly/colly"
)



func GetConcertBookPage() string {
	var url string = GetUrl()
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
