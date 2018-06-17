package main

import (
	"os"

	auth "github.com/aflashyrhetoric/payup-api/auth"
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

	authRoutes := router.Group("/api/v1/auth")
	{
		receiptRoutes.POST("/", auth.CreateUser)
		receiptRoutes.GET("/:id", auth.ShowUser)
		receiptRoutes.PUT("/:id", auth.UpdateUser)
		receiptRoutes.DELETE("/:id", auth.DeleteUser)
	}
	router.Run()
}
