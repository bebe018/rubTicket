package WD

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"server/concertInfo"
	"strconv"
	"github.com/tebeka/selenium"
)

var closest int = math.MaxInt32

func ChooseConcertEvent(wd selenium.WebDriver) {
	concertBookPage := ""
Here:
	if concertBookPage = concertInfo.GetConcertBookPage(); concertBookPage == "" {
		fmt.Println("can not find concert")
		goto Here
	}
	wd.Get(concertBookPage)

	settings := concertInfo.GetJSON()

	var element selenium.WebElement
	webElement, err := wd.FindElement(selenium.ByClassName, "zone")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover happened")
		}
	}()

	if elems, err := webElement.FindElements(selenium.ByClassName, "select_form_b"); len(elems) != 0 {
		if err != nil {
			log.Println(err)
		}
		element = chooseTheCloset(elems, settings.Price)
	} else if elems, err := wd.FindElements(selenium.ByClassName, "select_form_a"); len(elems) != 0 {
		if err != nil {
			log.Println(err)
		}
		element = elems[0]
	}
	element, err = element.FindElement(selenium.ByTagName, "a")
	if err != nil {
		log.Println(err)
	}
	err = element.Click()
	if err != nil {
		log.Println(err)
	}
}

func chooseTheCloset(elements []selenium.WebElement, price int) selenium.WebElement {
	element := elements[0]
	for _, elem := range elements {
		if extractPriceString, err := extractPrice(elem); err == nil {
			extractPriceInt, err := strconv.Atoi(extractPriceString)
			if err != nil {
				log.Println(err)
			}
			if int(math.Abs(float64(extractPriceInt)-float64(price))) < closest {
				closest = int(math.Abs(float64(extractPriceInt) - float64(price)))
				element = elem
			}
		}
	}
	return element
}

func extractPrice(elem selenium.WebElement) (string, error) {
	pattern := `\d+\D+(\d+)`
	re := regexp.MustCompile(pattern)
	text, err := elem.Text()
	if err != nil {
		log.Println(err)
	}
	matches := re.FindStringSubmatch(text)
	if matches != nil {
		return matches[1], nil
	}
	return "", err
}
