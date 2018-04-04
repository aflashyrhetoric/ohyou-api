package main

import (
	"os"

	e "github.com/aflashyrhetoric/payup-api/expense"
	r "github.com/aflashyrhetoric/payup-api/receipt"
	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("PORT", "8114")

	// Configure Router
	router := gin.Default()
	expenseRoutes := router.Group("/api/v1/expenses")
	{
		expenseRoutes.GET("", e.ListExpenses)
		expenseRoutes.POST("/", e.CreateExpense)
		expenseRoutes.GET("/:id", e.ShowExpense)
		expenseRoutes.PUT("/:id", e.UpdateExpense)
		expenseRoutes.DELETE("/:id", e.DeleteExpense)
	}

	receiptRoutes := router.Group("/api/v1/receipts")
	{
		// Disable ListReceipts
		// receiptRoutes.GET("/", r.ListReceipts)
		receiptRoutes.POST("/", r.CreateReceipt)
		receiptRoutes.GET("/:id", r.ShowReceipt)
		receiptRoutes.PUT("/:id", r.UpdateReceipt)
		receiptRoutes.DELETE("/:id", r.DeleteReceipt)
	}

	router.Run()
}
