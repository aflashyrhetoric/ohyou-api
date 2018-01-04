package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/transactions")
	{
		v1.POST("/", createTransaction)
		v1.GET("/", listTransaction)
		// v1.GET("/:id", showTransaction)
		// v1.PUT("/:id", updateTransaction)
		// v1.DELETE("/:id", deleteTransaction)
	}
	router.Run()
}

func init() {
	// Initialize sql.DB
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/sqltest")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
