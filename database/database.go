package database

import (
	"database/sql"
	"log"

	// Imported for testing purposes
	_ "github.com/go-sql-driver/mysql"
)

// var DatabaseInfo model.Database

// NewDB ...returns a pointer to the database
func NewDB() (*sql.DB, error) {
	// Initialize db
	var err error
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/payup_api")
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
