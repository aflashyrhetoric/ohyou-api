package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {

	// Initialize db
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost)/ohyou_api")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Configure Router
	router := gin.Default()
	v1 := router.Group("/api/v1/transactions")
	{
		v1.GET("", listTransactions)
		v1.POST("/", createTransaction)
		v1.GET("/:id", showTransaction)
		// v1.PUT("/:id", updateTransaction)
		// v1.DELETE("/:id", deleteTransaction)
	}
	router.Run()
}
