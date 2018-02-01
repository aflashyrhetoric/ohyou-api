package main

import (
	_ "github.com/go-sql-driver/mysql"

	m "github.com/aflashyrhetoric/payup-api/model"
	"github.com/gin-gonic/gin"
)

// var db *sql.DB

func main() {

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
