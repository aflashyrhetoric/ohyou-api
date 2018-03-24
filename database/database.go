package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Imported for testing purposes
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// var DatabaseInfo model.Database

// NewDB ...returns a pointer to the database
func NewDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBConnectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/payup_api", DBUser, DBPassword, DBHost)
	fmt.Println(DBConnectionString)
	// Initialize db
	db, err := sql.Open("mysql", DBConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
