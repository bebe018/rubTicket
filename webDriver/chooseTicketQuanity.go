package WD

import (
	"fmt"
	"log"
	"server/concertInfo"
	"github.com/tebeka/selenium"
)

func TicketQuantity(wd selenium.WebDriver) {

	settings := concertInfo.GetJSON()

	elems, err := wd.FindElements(selenium.ByTagName, "option")
	for _, elem := range elems {
		tickets, _ := elem.Text()
		if tickets == settings.WantTickets {
			err = elem.Click()
			if err != nil {
				log.Println(err)
			}
		}
	}
	elem, err := wd.FindElement(selenium.ByID, "TicketForm_agree")
	if err != nil {
		log.Println(err)
	}
	err = elem.Click()
	if err != nil {
		log.Println(err)
	}
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("ticketQuantity recover happened")
		}
	}()
}
