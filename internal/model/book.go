package model

type UserBook struct {
	Key           string   `json:"key"`
	Title         string   `json:"title"`
	Author        []string `json:"author"`
	EditionNumber int      `json:"edition_number"`
}
