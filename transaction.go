package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type (
	transaction struct {
		gorm.Model
		Description string `json:"description"`
		Purchaser   uint   `json:"purchaser"`
		Amount      uint   `json:"amount"`
	}
	transformedTransaction struct {
		ID            uint   `json:"id"`
		Description   string `json:"description"`
		Purchaser     uint   `json:"purchaser"`
		Amount        uint   `json:"amount"`
		Beneficiaries []User `json:"beneficiaries"`
		Users         []User `gorm:"many2many:user_beneficiaries;"`
	}
)

func createTransaction(c *gin.Context) {
	// description := json.Marshal(true)
	description := c.PostForm("description")
	amount, _ := ConvertDollarsToCents(strconv.ParseFloat(c.PostForm("amount"), 32))
	beneficiaries := c.PostForm("beneficiaries")
	transaction := transaction{Description: description, Amount: amount, Beneficiaries: beneficiaries}
	db.Save(&transaction)
	c.JSON(
		http.StatusCreated,
		gin.H{"status": http.StatusCreated, "message": "Transaction created successfully.", "resourceId": transaction.ID})
}

func listTransaction(c *gin.Context) {
	var transactions []transaction
	var _transactions []transformedTransaction

	db.Find(&transactions)

	if len(transactions) <= 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}

	for _, item := range transactions {
		_transactions = append(_transactions, transformedTransaction{ID: item.ID, Description: item.Description, Amount: item.Amount, Beneficiaries: item.Beneficiaries})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _transactions})
}

// func showTransaction(c *gin.Context) {
// 	var transaction transaction
// 	transactionID := c.Param("id")

// 	db.First(&transaction, transactionID)

// 	if transaction.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
// 		return
// 	}

// 	completed := false

// 	if transaction.Completed == 1 {
// 		completed = true
// 	} else {
// 		completed = false
// 	}

// 	_transaction := transformedTransaction{ID: transaction.ID, Title: transaction.Title, Completed: completed}

// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _Transaction})
// }

// func updateTransaction(c *gin.Context) {
// 	var transaction transaction
// 	transactionid := c.Param("id")

// 	db.First(&transaction, transactionID)

// 	if transaction.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
// 		return
// 	}

// 	db.Model(&transaction).Update("title", c.PostForm("title"))
// 	completed, _ := strconv.Atoi(c.PostForm("completed"))
// 	db.Model(&transaction).Update("completed", completed)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Transaction updated successfully!"})
// }

// func deleteTransaction(c *gin.Context) {
// 	var transaction transaction
// 	transactionid := c.Param("id")
// 	db.First(&transaction, transactionID)
// 	if transaction.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
// 		return
// 	}
// 	db.Delete(&transaction)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Transaction deleted successfully!"})
// }
