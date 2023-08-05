package scryfall

import "github.com/google/uuid"

type Card struct {
	Name string `json:"name"`
}

type CoreCard struct {
	ArenaId             int
	Id                  uuid.UUID
	Lang                string
	Mtgo_id             int
	Mtgo_foil_id        int
	Multiverse_ids      []int
	Tcgplayer_id        int
	Tcgplayer_etched_id int
	Cardmarket_id       int
	Object              string
	Oracle_id           uuid.UUID
	Prints_search_uri   string
	Rulings_uri         string
	Scryfall_uri        string
	Uri                 string
}
