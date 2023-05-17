package WD

import (
	"github.com/tebeka/selenium"
	"time"
	"encoding/json"
	"os"
	"io/ioutil"
	"fmt"
)


func getCookie() selenium.Cookie {
	jsonFile, err := os.Open("cookie.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var cookie selenium.Cookie
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cookie)
	return cookie
}

func Login() selenium.WebDriver {
	cookie := getCookie()
	WebDriver := GetWebDriver()
	safeGet(WebDriver, "https://tixcraft.com")
	time.Sleep(1 * time.Second)
	clickCookieButton(WebDriver)
	WebDriver.DeleteCookie("SID")
	WebDriver.AddCookie(&cookie)
	loginWithGoogle(WebDriver)
	return WebDriver
}
