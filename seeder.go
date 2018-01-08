package main

import (
	"log"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/malisit/kolpa"
)

// SeedTransactions ... Seeds database with sample data.
func SeedTransactions(c *gin.Context) {
	k := kolpa.C()

	numberOfRecords := 25
	groupCount := 4

	// Create n records
	for i := 0; i < numberOfRecords; i++ {
		description := k.LoremWord()
		purchaser := rand.Intn(groupCount) + 1
		amount := rand.Intn(8000) + 100
		// Build up model to be saved
		newTransaction := Transaction{
			description: description,
			purchaser:   int(purchaser),
			amount:      amount,
		}
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare("INSERT INTO transactions VALUES(NULL, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		stmt.Exec(
			newTransaction.description,
			newTransaction.purchaser,
			newTransaction.amount)
		// Save
		tx.Commit()
	}

}
