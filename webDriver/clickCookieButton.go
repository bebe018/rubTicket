package WD

import (
	"github.com/tebeka/selenium"
	"log"
)
func ClickCookieButton(wd selenium.WebDriver) {
	cookieButton, err := wd.FindElement(selenium.ByID, "onetrust-accept-btn-handler")
	if err != nil {
		panic(err)
	}
	err = cookieButton.Click()
	if err != nil {
		panic(err)
	}
	err = wd.MaximizeWindow("")
	if err != nil {
		log.Fatal(err)
	}
}