package models

type Article struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Anons  string `json:"anons"`
	Text   string `json:"text"`
	Photo  string `json:"photo"`
}
