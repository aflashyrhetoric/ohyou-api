package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/transactions")
	{
		v1.GET("/", listTransaction)
		// v1.POST("/", createTransaction)
		// v1.GET("/:id", showTransaction)
		// v1.PUT("/:id", updateTransaction)
		// v1.DELETE("/:id", deleteTransaction)
	}
	router.Run()
}

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:***REMOVED***@/godo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&TransactionModel{})
}

type (
	// TransactionModel describes a TransactionModel type
	TransactionModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}

	transformedTransaction struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

func createTransaction(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	Transaction := TransactionModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&Transaction)
	c.JSON(
		http.StatusCreated,
		gin.H{"status": http.StatusCreated, "message": "Transaction created successfully.", "resourceId": Transaction.ID})
}

func listTransaction(c *gin.Context) {
	var Transactions []TransactionModel
	var _Transactions []transformedTransaction

	db.Find(&Transactions)
	if len(Transactions) <= 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}

	for _, item := range Transactions {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_Transactions = append(_Transactions, transformedTransaction{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _Transactions})
}

func showTransaction(c *gin.Context) {
	var Transaction TransactionModel
	TransactionID := c.Param("id")

	db.First(&Transaction, TransactionID)

	if Transaction.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}

	completed := false

	if Transaction.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_Transaction := transformedTransaction{ID: Transaction.ID, Title: Transaction.Title, Completed: completed}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _Transaction})
}

func updateTransaction(c *gin.Context) {
	var Transaction TransactionModel
	TransactionID := c.Param("id")

	db.First(&Transaction, TransactionID)

	if Transaction.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}

	db.Model(&Transaction).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&Transaction).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Transaction updated successfully!"})
}

func deleteTransaction(c *gin.Context) {
	var Transaction TransactionModel
	TransactionID := c.Param("id")
	db.First(&Transaction, TransactionID)
	if Transaction.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}
	db.Delete(&Transaction)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Transaction deleted successfully!"})
}
