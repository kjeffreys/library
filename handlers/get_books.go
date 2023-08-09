// handlers/get_books.go

package handlers

import (
	"encoding/json"
	"library/database"
	"library/models"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

	rows, err := db.Query("SELECT id, title, author, publishedyear, genre, summary FROM books")
	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.Genre, &book.Summary); err != nil {
			http.Error(w, "Failed to retrieve book", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}
