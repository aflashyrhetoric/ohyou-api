package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	// "github.com/aflashyrhetoric/ohyou-api/database"
	m "github.com/aflashyrhetoric/ohyou-api/model"
	"github.com/aflashyrhetoric/ohyou-api/utils"
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
		v1.GET("", m.listTransactions)
		v1.POST("/", m.createTransaction)
		v1.GET("/:id", m.showTransaction)
		v1.PUT("/:id", m.updateTransaction)
		v1.DELETE("/:id", m.deleteTransaction)
	}
	router.Run()
}
