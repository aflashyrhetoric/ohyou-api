package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func seedTransactions(c *gin.Context) {
	// Retrieve POST values
	description := c.PostForm("description")
	purchaser, err := strconv.ParseInt(c.PostForm("purchaser"), 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	amount, err := ConvertDollarsStringToCents(c.PostForm("amount"))
	if err != nil {
		log.Fatal(err)
	}
	// Build up model to be saved
	newTransaction := transformedTransaction{
		Description: description,
		Purchaser:   int(purchaser),
		Amount:      amount,
	}
	// Save
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO transactions VALUES(NULL, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(
		newTransaction.Description,
		newTransaction.Amount,
		newTransaction.Purchaser)
	tx.Commit()

}
