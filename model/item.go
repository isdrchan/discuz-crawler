package model

type Item struct {
	Id      string `json:"id"`
	Url     string `json:"url"`
	Section string `json:"section"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
