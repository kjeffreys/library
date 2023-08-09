// main.go

package main

import (
	"library/database"
	"library/handlers"
	"library/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Connect to the database
	database.Connect()
	// Initialize the database (e.g., create tables if they don't exist)
	database.InitializeDatabase()

	// can extract to router.go fn, then r = NewRouter(), any new router logic
	r := mux.NewRouter()

	r.Use(middlewares.LoggingMiddleware)

	// Define our routes
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", handlers.DeleteBook).Methods("DELETE")

	// Start the HTTP server
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
