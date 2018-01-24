package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	m "github.com/aflashyrhetoric/ohyou-api/model"
)

// var db *sql.DB

func main() {

	// Initialize db
	var err error
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/ohyou_api")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Configure Router
	router := gin.Default()
	v1 := router.Group("/api/v1/transactions")
	{
		v1.GET("", m.ListTransactions)
		v1.POST("/", m.CreateTransaction)
		v1.GET("/:id", m.ShowTransaction)
		v1.PUT("/:id", m.UpdateTransaction)
		v1.DELETE("/:id", m.DeleteTransaction)
	}
	router.Run()
}
