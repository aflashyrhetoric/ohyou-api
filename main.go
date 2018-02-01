package main

import (
	"database/sql"

	"github.com/aflashyrhetoric/payup-api/database"
	t "github.com/aflashyrhetoric/payup-api/transaction"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// DB ... points to the database
var DB *sql.DB

func main() {

	DB = database.NewDB()

	// Configure Router
	router := gin.Default()
	v1 := router.Group("/api/v1/transactions")
	{
		v1.GET("", t.ListTransactions)
		v1.POST("/", t.CreateTransaction)
		v1.GET("/:id", t.ShowTransaction)
		v1.PUT("/:id", t.UpdateTransaction)
		v1.DELETE("/:id", t.DeleteTransaction)
	}
	router.Run()
}
