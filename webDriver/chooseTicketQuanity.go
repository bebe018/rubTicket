package WD

import (
	"github.com/tebeka/selenium"
)


func TicketQuantity(wd selenium.WebDriver, wantTickets string) {
	// elem, err := wd.FindElement(selenium.ByTagName, "select")
	elems, err := wd.FindElements(selenium.ByTagName, "option")
	for _, elem := range elems {
		tickets, _ := elem.Text()
		if tickets == wantTickets {
			err = elem.Click()
			check(err)
		}
	}
	elem, err := wd.FindElement(selenium.ByID, "TicketForm_agree")
	check(err)
	err = elem.Click()
	check(err)
}