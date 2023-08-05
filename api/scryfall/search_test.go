package scryfall

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	cards := Search("meren")

	fmt.Printf("Cards count: %d", len(cards))
}
