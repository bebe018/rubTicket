package WD

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"github.com/tebeka/selenium"
)

var closest int = math.MaxInt32

func ChooseConcertEvent(wd selenium.WebDriver, price int) {

	elems, err := wd.FindElements(selenium.ByClassName, "select_form_b")
	check(err)

	value := "a"
	element := elems[0]
	for _, elem := range elems {
		extractPriceString := extractPrice(elem)
		extractPriceInt, err := strconv.Atoi(extractPriceString)
		check(err)
		if int(math.Abs(float64(extractPriceInt)-float64(price))) < closest {
			closest = int(math.Abs(float64(extractPriceInt) - float64(price)))
			element = elem
		}
	}
	element, err = element.FindElement(selenium.ByTagName, value)
	check(err)
	err = element.Click()
	check(err)
}

func extractPrice(elem selenium.WebElement) string {
	pattern := `\d+\D+(\d+)`
	re := regexp.MustCompile(pattern)
	text, err := elem.Text()
	check(err)
	matches := re.FindStringSubmatch(text)
	if matches != nil {
		return matches[1]
	}
	return ""
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
