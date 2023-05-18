package WD

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const port = 8080

func GetWebDriver() selenium.WebDriver {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36",
		},
	}
	caps.AddChrome(chromeCaps)
	_, err := selenium.NewChromeDriverService("chromedriver", port, opts...)
	if err != nil {
		panic(err)
	}
	var cap_extension chrome.Capabilities
	cap_extension.AddExtension("modheader.crx")
	caps.AddChrome(cap_extension)

	WebDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	WebDriver.Get("https://webdriver.modheader.com/add?test=ModHeader%20Test")
	
	return WebDriver
}
