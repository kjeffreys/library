// models/book.go

package models

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"publishedYear"`
	Genre         string `json:"genre"`
	Summary       string `json:"summary"`
}
