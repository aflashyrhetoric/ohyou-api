package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	transaction struct {
		description string
		purchaser   int
		amount      int
	}
	transformedTransaction struct {
		ID            int    `json:"id"`
		Description   string `json:"description"`
		Purchaser     int    `json:"purchaser"`
		Amount        int    `json:"amount"`
		Beneficiaries []User `json:"beneficiaries"`
	}
)

func createTransaction(c *gin.Context) {
	// Retrieve POST values
	description := c.PostForm("description")

	amount, err := ConvertDollarsStringToCents(c.PostForm("amount"))
	if err != nil {
		log.Fatal(err)
	}

	purchaser, err := strconv.ParseInt(c.PostForm("purchaser"), 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Build up model
	newTransaction := transformedTransaction{
		ID:            3,
		Description:   description,
		Purchaser:     int(purchaser),
		Amount:        amount,
		Beneficiaries: nil,
	}

	// Save
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO transactions VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(
		newTransaction.ID,
		newTransaction.Description,
		newTransaction.Amount,
		newTransaction.Purchaser)
	tx.Commit()

	// Response
	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":     http.StatusCreated,
			"message":    "Transaction created successfully.",
			"resourceId": newTransaction.ID,
		},
	)
}

func listTransactions(c *gin.Context) {
	var (
		ID           int
		Description  string
		Purchaser    int
		Amount       int
		responseData []transformedTransaction
	)

	// Save
	stmt, err := db.Prepare("SELECT * FROM transactions")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ID, &Description, &Purchaser, &Amount)
		if err != nil {
			log.Fatal(err)
		}
		responseData = append(responseData, transformedTransaction{ID, Description, Purchaser, Amount, nil})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responseData})
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

// 	_transaction := TransformedTransaction{ID: transaction.ID, Title: transaction.Title, Completed: completed}

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
