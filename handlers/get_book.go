// handlers/get_book.go

package handlers

import (
	"encoding/json"
	"library/database"
	"library/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

	vars := mux.Vars(r)
	id := vars["id"]

	var book models.Book
	sqlStatement := `SELECT id, title, author, publishedyear, genre, summary FROM books WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.Genre, &book.Summary); err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}
