package concertInfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"server/search"
	"strings"
	"time"
)

type Settings struct {
	Query       string `json:"query"`
	Date        string `json:"date"`
	Price       int    `json:"price"`
	WantTickets string `json:"wantTickets"`
}

func GetJSON() Settings {
	jsonFile, err := os.Open("settings.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var settings Settings
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &settings)
	return settings
}

// func getConcertPage() string {
// 	settings := GetJSON()
// 	setConcertInfo()
// 	concertInfo := getConcertInfo()
// 	concertPage := search.FuzzySearch(settings.Query, concertInfo)
// 	concertPage = strings.ReplaceAll(concertPage, "detail", "game")
// 	return concertPage
// }

// func GetUrl() string {
// 	concertPage := getConcertPage()
// 	for concertPage == "" {
// 		time.Sleep(1 * time.Second)
// 		concertPage = getConcertPage()
// 	}
// 	url := fmt.Sprintf("https://tixcraft.com%s", concertPage)
// 	return url
// }

func GetUrl() string {
	settings := GetJSON()
	setConcertInfo()
	concertInfo := getConcertInfo()
	concertPage := ""
Here:
	if concertPage = search.FuzzySearch(settings.Query, concertInfo); concertPage == "" {
		time.Sleep(1 * time.Second)
		goto Here
	}
	concertPage = search.FuzzySearch(settings.Query, concertInfo)
	concertPage = strings.ReplaceAll(concertPage, "detail", "game")
	url := fmt.Sprintf("https://tixcraft.com%s", concertPage)
	return url
}
