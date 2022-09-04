package models

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Year        string `json:"year"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}
