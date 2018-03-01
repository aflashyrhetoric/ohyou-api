package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Imported for testing purposes
	_ "github.com/go-sql-driver/mysql"
)

// var DatabaseInfo model.Database

// NewDB ...returns a pointer to the database
func NewDB() (*sql.DB, error) {

	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBConnectionString := fmt.Sprint("%s:%s@tcp(%s:3306)/payup_api", DBUser, DBPassword, DBHost)
	fmt.Println(DBConnectionString)
	// Initialize db
	var err error
	db, err := sql.Open("mysql", DBConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
