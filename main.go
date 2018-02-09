package main

import (
	"os"

	t "github.com/aflashyrhetoric/payup-api/transaction"
	"github.com/gin-gonic/gin"
)

func main() {

  os.Setenv("PORT", "8114")


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
