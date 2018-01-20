package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func seedTransactions(c *gin.Context) {
	// Retrieve POST values
	description := c.PostForm("description")
	purchaser, err := getPurchaser(c)
	if err != nil {
		log.Fatal(err)
	}
	amount, err := getAmount(c)
	if err != nil {
		log.Fatal(err)
	}
	beneficiaries := loadTransactionBeneficiaryData()
	// Build up model to be saved
	newTransaction := transformedTransaction{
		Description:   description,
		Purchaser:     purchaser,
		Amount:        amount,
		Beneficiaries: beneficiaries,
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
		newTransaction.Purchaser,
		newTransaction.Beneficiaries)
	tx.Commit()

}
