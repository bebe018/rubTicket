package main

import (
	"server/webDriver"
	"time"
)

func main() {
	WebDriver := WD.Login()
	WD.ChooseConcertEvent(WebDriver)
	WD.TicketQuantity(WebDriver)
	defer WebDriver.Quit()
	time.Sleep(10 * time.Minute)
}
