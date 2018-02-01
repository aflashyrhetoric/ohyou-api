package database

import (
	"database/sql"
	"log"

	"github.com/aflashyrhetoric/payup-api/model"
	_ "github.com/go-sql-driver/mysql"
)

var DatabaseInfo model.Database

// New...returns a pointer to the database
func NewDB() *sql.DB {
	// Initialize db
	var err error
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/ohyou_api")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
