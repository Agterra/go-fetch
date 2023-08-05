package scryfall

type Response struct {
	Object      string `json:"object"`
	Total_cards int    `json:"total_cards"`
	Has_more    bool   `json:"has_more"`
	Data        []Card `json:"data"`
}
