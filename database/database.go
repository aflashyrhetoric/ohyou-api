package database

import (
	"database/sql"
)

var DatabaseInfo model.Database

func NewDB() *sql.DB {
	// Initialize db
	var err error
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/ohyou_api")
	if err != nil {
		log.Fatal(err)
  }
  
  return db
}

func InitDatabase() {
  db :=
}