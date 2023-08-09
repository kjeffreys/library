// handlers/delete_book.go

package handlers

import (
	"library/database"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

	vars := mux.Vars(r)
	id := vars["id"]

	sqlStatement := `DELETE FROM books WHERE id=$1`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
