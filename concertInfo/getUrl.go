package concertInfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"server/search"
	"strings"
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

func GetUrl() string {
	settings := GetJSON()
	setConcertInfo()
	concertInfo := getConcertInfo()
	concertPage := ""
Here:
	if concertPage = search.FuzzySearch(settings.Query, concertInfo); concertPage == "" {
		fmt.Println("can not find the concert")
		goto Here
	}
	concertPage = search.FuzzySearch(settings.Query, concertInfo)
	concertPage = strings.ReplaceAll(concertPage, "detail", "game")
	url := fmt.Sprintf("https://tixcraft.com%s", concertPage)
	return url
}
