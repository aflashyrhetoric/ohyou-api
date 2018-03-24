package main

import (
	"os"

	t "github.com/aflashyrhetoric/payup-api/expense"
	"github.com/gin-gonic/gin"
)

func main() {

  os.Setenv("PORT", "8114")


	// Configure Router
	router := gin.Default()
	v1 := router.Group("/api/v1/expenses")
	{
		v1.GET("", t.ListExpenses)
		v1.POST("/", t.CreateExpense)
		v1.GET("/:id", t.ShowExpense)
		v1.PUT("/:id", t.UpdateExpense)
		v1.DELETE("/:id", t.DeleteExpense)
	}
	router.Run()
}
