package WD

import (
	"github.com/tebeka/selenium"
)
func clickCookieButton(wd selenium.WebDriver) {
	cookieButton, err := wd.FindElement(selenium.ByID, "onetrust-accept-btn-handler")
	if err != nil {
		panic(err)
	}
	err = cookieButton.Click()
	if err != nil {
		panic(err)
	}
	
}