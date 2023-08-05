package scryfall

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var scryfallURL = "https://api.scryfall.com/cards/search?q="

/*
query The scryfall full text search, like 'o:"draw" c:"u" t:"creature'
*/
func Search(query string) []Card {
	resp, err := http.Get(fmt.Sprintf("%s%s", scryfallURL, query))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Print(string(body))

	var response Response

	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Fatal(err)
	}

	return response.Data
}
