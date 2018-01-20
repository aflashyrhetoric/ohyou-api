package main

import (
	"log"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/malisit/kolpa"
)

// Generates a random []int representing beneficiary IDs
func generateRandomBeneficiaries(maxNum int) []int {
	var randomBeneficiariesArr []int

	// Gens a random # in how MANY beneficiaries there are
	randomLoopCount := rand.Intn(maxNum) + 1

	// Ensure that randomBeneficiariesArr isn't empty
	for len(randomBeneficiariesArr) == 0 {
		for i := 0; i < randomLoopCount; i++ {
			randomBeneficiaryID := rand.Intn(maxNum)
			// If randomly generated beneficiary array doesn't already contain the ID, add it
			if !ArrayContainsInt(randomBeneficiaryID, randomBeneficiariesArr) {
				randomBeneficiariesArr = append(randomBeneficiariesArr, randomBeneficiaryID)
			}
		}
	}

	return randomBeneficiariesArr
}

// SeedTransactions ... Seeds database with sample data.
func SeedTransactions(c *gin.Context) {
	k := kolpa.C()

	numberOfRecords := 25
	groupCount := 4

	// Create numberOfRecords records
	for i := 0; i < numberOfRecords; i++ {
		description := k.LoremWord()
		purchaser := rand.Intn(groupCount) + 1
		amount := rand.Intn(8000) + 100

		// Step 1: Initial transaction
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare(`
			INSERT INTO transactions 
			VALUES(NULL, ?, ?, ?)
		`)
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(
			description,
			purchaser,
			amount)

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		// Save
		tx.Commit()

		// Step 2: Generate TransactionBeneficiaries data
		beneficiaries := generateRandomBeneficiaries(groupCount)

		for _, beneficiaryID := range beneficiaries {
			tx, err = db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err = tx.Prepare(`
				INSERT INTO transactions_beneficiaries 
				VALUES(?, ?) 
			`)
			stmt.Exec(lastInsertID, beneficiaryID)
			if err != nil {
				log.Fatal(err)
			}
			tx.Commit()
		}
	}
}
