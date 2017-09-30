package main

import (
	"fmt"
	"os"
	// "bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/joho/godotenv"
)

type QiitaItems []struct {
	ID           string `json:"id"`
	URL          string `json:"url"`
	Private      bool   `json:"private"`
	Coediting    bool   `json:"coediting"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	RenderedBody string `json:"rendered_body"`
}

func main() {
	dotenverr := godotenv.Load()
	if dotenverr != nil {
		log.Fatal("Error loading .env file")
	}
	request, _ := http.NewRequest("GET", "https://qiita.com/api/v2/authenticated_user/items?per_page=100", nil)
	token := os.Getenv("QIITA_ACCESS_TOKEN")
	request.Header.Set("Authorization", "Bearer " + token)

	client := new(http.Client)
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error")
			log.Fatal(err)
		}
		var results QiitaItems
		json.Unmarshal(body, &results)
		fmt.Println(results)
	}
}

func getQiitaItems() {

}
