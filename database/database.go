package database

import (
	"database/sql"
	"log"

	"github.com/aflashyrhetoric/payup-api/model"
)

var DatabaseInfo model.Database

// New...returns a pointer to the database
func New() *sql.DB {
	// Initialize db
	var err error
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/ohyou_api")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func init() {
	db := New()
	return db
}
