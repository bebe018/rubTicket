package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"math"
	"server/concertInfo"
	"server/search"
	"server/webDriver"
	"strings"
	"time"
)

var cookie = selenium.Cookie{
	Name:   "SID",
	Value:  "mhmu2k9kc9jr192deh6444lqc1", // 需要定期更換
	Domain: "tixcraft.com",
	Path:   "/",
	Secure: true,
	Expiry: math.MaxUint32,
}

const query string = "建宮蓋廟" // 如有英文請輸入全小寫
const date string = "2023/07/16"
const price int = 1500
const wantTickets = "2"

func main() {
	WebDriver := WD.GetWebDriver()
	WD.ClickCookieButton(WebDriver)
	WebDriver.DeleteCookie("SID")
	WebDriver.AddCookie(&cookie)
	WD.LoginWithGoogle(WebDriver)

	concertInfo.SetConcertInfo()
	concertInfo := concertInfo.GetConcertInfo()
	concertPage := search.FuzzySearch(query, concertInfo)
	concertPage = strings.ReplaceAll(concertPage, "detail", "game")

	url := fmt.Sprintf("https://tixcraft.com%s", concertPage)

	concertBookPageMap := search.GetConcertBookPage(url)
	concertBookPage := concertBookPageMap[date]
	WebDriver.Get(concertBookPage)
	WD.ChooseConcertEvent(WebDriver, price)
	WD.TicketQuantity(WebDriver, wantTickets)

	defer WebDriver.Quit()
	time.Sleep(10 * time.Minute)
}
