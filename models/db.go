package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		print(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	content, err := os.ReadFile("./database/init.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %s", err)
	}
	if _, err := db.Exec(string(content)); err != nil {
		log.Fatalf("failed to execute SQL commands: %s", err)
	}
	log.Println("Database initialized successfully.")
}
