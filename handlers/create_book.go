// handlers/create_book.go

package handlers

import (
	"encoding/json"
	"library/database"
	"library/models"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

	var book models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	sqlStatement := `
		INSERT INTO books (title, author, publishedyear, genre, summary)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, book.Title, book.Author, book.PublishedYear, book.Genre, book.Summary).Scan(&id)
	if err != nil {
		http.Error(w, "Failed to insert book", http.StatusInternalServerError)
		return
	}

	book.ID = id
	json.NewEncoder(w).Encode(book)
}
