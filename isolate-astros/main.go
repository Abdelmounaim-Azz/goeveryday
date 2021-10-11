package main

import (
	"encoding/json"
	"fmt"

	"github.com/Abdelmounaim-Azz/isolate-astros/astros"
)

func getAstronauts(getWebRequest astros.GetWebRequest) (int, error) {
	url := "http://api.open-notify.org/astros.json"
	body, err := getWebRequest.FetchBytes(url)
	if err != nil {
		return 0, err
	}
	peopleResult := astros.People{}
	if err := json.Unmarshal(body, &peopleResult); err != nil {
		return 0, err
	}
	return peopleResult.Number, err
}
func main() {
	liveClient := astros.LiveGetWebRequest{}
	number, err := getAstronauts(liveClient)
	if err != nil {
		panic(err)
	}
	fmt.Println(number)
}
