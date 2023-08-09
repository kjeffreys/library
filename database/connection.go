// database/connection.go

package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() *sql.DB {
	if db != nil {
		return db
	}

	connStr := os.Getenv("DATABASE_URL") // get the connection string from an environment variable
	if connStr == "" {
		log.Fatal("Database connection string is not set.")
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitializeDatabase() {
	// Initialization script for the books table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author VARCHAR(255) NOT NULL,
        publishedyear INT NOT NULL,
        genre VARCHAR(255),
        summary TEXT
    );`)
	if err != nil {
		log.Fatalf("Failed to execute initialization script: %v", err)
	}
}
