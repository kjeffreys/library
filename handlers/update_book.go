// handlers/update_book.go

package handlers

import (
	"encoding/json"
	"library/database"
	"library/models"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

	vars := mux.Vars(r)
	id := vars["id"]

	var book models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	sqlStatement := `
		UPDATE books
		SET title=$1, author=$2, publishedyear=$3, genre=$4, summary=$5
		WHERE id=$6`
	_, err := db.Exec(sqlStatement, book.Title, book.Author, book.PublishedYear, book.Genre, book.Summary, id)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}
