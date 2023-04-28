package concertdata

import (
	// fuzzy "server/search"
	"github.com/gocolly/colly"
)

var concertInfo = make(map[string]string)

func SetConcertInfo() {
	c := colly.NewCollector()
	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	c.OnHTML(".thumbnails", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
		e.ForEach("a[href]", func(i int, h *colly.HTMLElement) {
			concertName := h.ChildText("div[class='data'] > div[class='multi_ellipsis']")
			link := h.Attr("href")
			concertInfo[concertName] = link
		})

	})
	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})
	c.Visit("https://tixcraft.com/activity")
}

func GetConcertInfo() map[string]string{
	return concertInfo
}
