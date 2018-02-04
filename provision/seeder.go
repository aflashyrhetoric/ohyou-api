package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/aflashyrhetoric/payup-api/database"
	"github.com/aflashyrhetoric/payup-api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/malisit/kolpa"
)

// SeedTransactions ... Seeds database with sample data.
func main() {

	// Connect to database
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	k := kolpa.C()

	numberOfRecords := 25
	groupCount := 4

	// Create numberOfRecords records
	for i := 0; i < numberOfRecords; i++ {
		description := k.LoremParagraph()
		description = description[:25]
		purchaser := rand.Intn(groupCount) + 1
		amount := rand.Intn(8000) + 100

		// Step 1: Initial transaction
		fmt.Printf("Seeding Transaction %v...\n", i+1)
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
		fmt.Print("...initial transaction saved!\n")

		// Step 2: Generate TransactionBeneficiaries data
		beneficiaries := generateRandomBeneficiaries(groupCount)

		for index, beneficiaryID := range beneficiaries {
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
			fmt.Printf("...beneficiary %v saved!\n", index)
		}
		fmt.Print("...success!\n")
		fmt.Print("-----------\n")
	}
}

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
			if !utils.ArrayContainsInt(randomBeneficiaryID, randomBeneficiariesArr) {
				randomBeneficiariesArr = append(randomBeneficiariesArr, randomBeneficiaryID)
			}
		}
	}

	// return []int{1, 2, 3}
	return randomBeneficiariesArr
}
