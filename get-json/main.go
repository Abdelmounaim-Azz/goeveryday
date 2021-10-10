package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type person struct {
	Name string `json:"name"`
}

type people struct {
	Number int      `json:"number"`
	Person []person `json:"people"`
}

func main() {
	apiURL := "http://api.open-notify.org/astros.json"
	people, err := getAstros(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d people found in space.\n", people.Number)
	for _, p := range people.Person {
		fmt.Printf("Let's wave to: %s\n", p.Name)
	}
}
func getAstros(apiURL string) (people, error) {
	p := people{}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return p, err
	}
	req.Header.Set("User-Agent", "spacecount")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return p, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(body, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}
