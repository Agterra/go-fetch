package main

import (
	"encoding/json"
	"fetch/v1/api/scryfall"
	"fmt"
)

func main() {
	var input string

	fmt.Print("Enter a scryfall query: ")
	fmt.Scan(&input)

	cards := scryfall.Search(input)

	str, _ := json.MarshalIndent(cards, "", "  ")
	fmt.Println("", string(str))
}
