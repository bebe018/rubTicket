package concertInfo

import (
	"github.com/gocolly/colly"
)

var concertInfo = make(map[string]string)

func setConcertInfo() {
	c := colly.NewCollector()
	
	c.OnHTML(".thumbnails", func(e *colly.HTMLElement) { 
		e.ForEach("a[href]", func(i int, h *colly.HTMLElement) {
			concertName := h.ChildText("div[class='data'] > div[class='multi_ellipsis']")
			link := h.Attr("href")
			concertInfo[concertName] = link
		})

	})
	c.OnRequest(func(r *colly.Request) { 
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})
	c.Visit("https://tixcraft.com/activity")
}

func getConcertInfo() map[string]string{
	return concertInfo
}
