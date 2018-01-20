package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	transaction struct {
		id          int
		description string
		purchaser   int
		amount      int
	}
	transformedTransaction struct {
		ID            int    `json:"id"`
		Description   string `json:"description"`
		Purchaser     int    `json:"purchaser"`
		Amount        int    `json:"amount"`
		Beneficiaries []int  `json:"beneficiaries"`
	}
)

// Get beneficiaries from POST variable
func getBeneficiaries(c *gin.Context) []int {

	// String IDs
	rawBeneficiaryIDs := c.PostForm("beneficiaries")

	// Array of string IDs
	beneficiaryStringIDs := strings.Split(rawBeneficiaryIDs, ",")

	var beneficiaries []int

	// Convert array of String to array of Int
	for _, element := range beneficiaryStringIDs {
		stringIDConvertedToInteger, err := strconv.Atoi(element)
		if err != nil {
			log.Print(err)
		}
		beneficiaries = append(beneficiaries, stringIDConvertedToInteger)
	}

	return beneficiaries
}

// Load beneficiaries from database
func loadTransactionBeneficiaryData(transactionID int) []int {
	var (
		beneficiaryIDs []int
	)
	stmt, err := db.Prepare(`
		SELECT * FROM transactions 
		WHERE transaction_id = ?
	`)
	if err != nil {
		log.Print(err)
	}
	rows, err := stmt.Query(transactionID)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	// Scan values to Go variables
	for rows.Next() {
		err := rows.Scan(&transactionID, beneficiaryIDs)
		if err != nil {
			log.Print(err)
		}
		beneficiaryIDs = append(beneficiaryIDs)
	}
	return beneficiaryIDs
}

func getID(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param("id"))
}

func getDescription(c *gin.Context) string {
	return c.PostForm("description")
}

func getPurchaser(c *gin.Context) (int, error) {
	return strconv.Atoi(c.PostForm("purchaser"))
}

func getAmount(c *gin.Context) (int, error) {
	return ConvertDollarsStringToCents(c.PostForm("amount"))
}

func createTransaction(c *gin.Context) {

	// Initial transaction

	// Retrieve POST values
	description := getDescription(c)
	purchaser, err := getPurchaser(c)
	if err != nil {
		log.Print(err)
	}
	amount, err := getAmount(c)
	if err != nil {
		log.Print(err)
	}
	beneficiaries := getBeneficiaries(c)

	// Build up model to be saved
	newTransaction := transformedTransaction{
		Description:   description,
		Purchaser:     int(purchaser),
		Amount:        amount,
		Beneficiaries: beneficiaries,
	}

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO transactions 
		VALUES(NULL, ?, ?, ?)
	`)
	if err != nil {
		log.Print(err)
	}
	res, err := stmt.Exec(
		newTransaction.Description,
		newTransaction.Amount,
		newTransaction.Purchaser)
	if err != nil {
		log.Print(err)
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	// Save beneficiaries data

	// Save a transaction (to the junction table) for each beneficiary
	for _, beneficiaryID := range newTransaction.Beneficiaries {
		tx, err = db.Begin()
		if err != nil {
			log.Print(err)
		}
		stmt, err = tx.Prepare(`
			INSERT INTO transactions_beneficiaries 
			VALUES(?, ?) 
		`)
		stmt.Exec(
			lastInsertID,
			beneficiaryID)
		if err != nil {
			log.Print(err)
		}
		tx.Commit()
	}

	// Response
	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Transaction created successfully.",
			"data":    lastInsertID,
		},
	)
}

func listTransactions(c *gin.Context) {
	var (
		ID            int
		Description   string
		Purchaser     int
		Amount        int
		BeneficiaryID int
		Beneficiaries []int
		responseData  []transformedTransaction
	)
	// Prepare SELECT statement
	stmt, err := db.Prepare("SELECT * FROM transactions")
	if err != nil {
		log.Print(err)
	}
	// Run Query
	rows, err := stmt.Query()
	if err != nil {
		log.Print(err)
	}

	defer rows.Close()
	// Scan values to Go variables
	for rows.Next() {
		err := rows.Scan(&ID, &Description, &Purchaser, &Amount)
		if err != nil {
			log.Print(err)
		}

		// Run second query to retrieve transactions_beneficiaries data
		stmt, err = db.Prepare(`
			SELECT beneficiary_id 
			FROM transactions_beneficiaries 
			WHERE transaction_id = ?
		`)
		if err != nil {
			log.Print(err)
		}
		// Run Query
		benRows, err := stmt.Query(ID)
		if err != nil {
			log.Print(err)
		}
		defer benRows.Close()
		// Scan values to Go variables
		for benRows.Next() {
			err := benRows.Scan(&BeneficiaryID)
			if err != nil {
				log.Print(err)
			}
			Beneficiaries = append(Beneficiaries, BeneficiaryID)
		}

		responseData = append(responseData, transformedTransaction{ID, Description, Purchaser, Amount, Beneficiaries})
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responseData})
}

func showTransaction(c *gin.Context) {
	var (
		ID            int
		Description   string
		Purchaser     int
		Amount        int
		BeneficiaryID int
		Beneficiaries []int
		responseData  transformedTransaction
	)
	transactionID, err := getID(c)
	if err != nil {
		log.Print(err)
	}
	// Check for invalid ID
	if transactionID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}
	// Prepare SELECT statement
	stmt, err := db.Prepare(`
		SELECT id, description, purchaser, amount
		FROM transactions
		WHERE id=?
	`)
	if err != nil {
		log.Print(err)
	}

	// Run Query
	row := stmt.QueryRow(transactionID)
	if err != nil {
		log.Print(err)
	}

	// Scan values to Go variables
	err = row.Scan(&ID, &Description, &Purchaser, &Amount)
	if err == sql.ErrNoRows {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Record not found",
			})
		return
	} else if err != nil {
		log.Print(err)
	}

	// Retrieve transactions_beneficiaries data
	stmt, err = db.Prepare(`
		SELECT beneficiary_id 
		FROM transactions_beneficiaries 
		WHERE transaction_id = ?
	`)
	if err != nil {
		log.Print(err)
	}
	// Run Query
	benRows, err := stmt.Query(ID)
	if err != nil {
		log.Print(err)
	}
	defer benRows.Close()
	// Scan values to Go variables
	for benRows.Next() {
		err := benRows.Scan(&BeneficiaryID)
		if err != nil {
			log.Print(err)
		}
		Beneficiaries = append(Beneficiaries, BeneficiaryID)
	}

	responseData = transformedTransaction{ID, Description, Purchaser, Amount, Beneficiaries}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responseData})
}

func updateTransaction(c *gin.Context) {

	// Retrieve POST update data
	transactionID, _ := getID(c)
	description := getDescription(c)
	purchaser, _ := getPurchaser(c)
	amount, _ := getAmount(c)

	// Check for invalid ID
	if transactionID == 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No Transaction found!",
			})
		return
	}
	// Prepare SELECT statement
	tx, err := db.Begin()
	stmt, err := tx.Prepare(`
		UPDATE transactions 
		SET description=?, purchaser=?, amount=?
		WHERE id=?;
	`)
	if err != nil {
		log.Print(err)
	}

	// Execute update query for regular transaction data
	_, err = stmt.Exec(description, purchaser, amount, transactionID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	// Update transactions_beneficiaries data
	beneficiaries := getBeneficiaries(c)

	// First, delete all old associated beneficiaries
	tx, err = db.Begin()
	stmt, err = tx.Prepare(`
		DELETE FROM transactions_beneficiaries 
		WHERE transaction_id=?;
	`)
	if err != nil {
		log.Print(err)
	}
	// Execute update query for beneficiaries
	_, err = stmt.Exec(transactionID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	// Second, insert new values
	for _, beneficiaryID := range beneficiaries {
		tx, err = db.Begin()
		stmt, err = tx.Prepare(`
			INSERT INTO transactions_beneficiaries 
			VALUES (?, ?);
		`)
		if err != nil {
			log.Print(err)
		}

		// Execute update query for beneficiaries
		_, err = stmt.Exec(transactionID, beneficiaryID)
		if err != nil {
			log.Print(err)
		}
		tx.Commit()
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Transaction updated successfully.",
		},
	)
}

func deleteTransaction(c *gin.Context) {
	transactionID, err := getID(c)
	if err != nil {
		log.Print(err)
	}
	// Check for invalid ID
	if transactionID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}
	// Prepare SELECT statement
	tx, err := db.Begin()
	stmt, err := tx.Prepare(`
		DELETE FROM transactions
		WHERE id=?
	`)
	if err != nil {
		log.Print(err)
	}

	// Run Query
	_, err = stmt.Exec(transactionID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	// Delete beneficiary data
	tx, err = db.Begin()
	stmt, err = tx.Prepare(`
		DELETE FROM transactions_beneficiaries
		WHERE transaction_id=?
	`)
	if err != nil {
		log.Print(err)
	}

	// Run Query
	_, err = stmt.Exec(transactionID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	responseMsg := fmt.Sprintf("Transaction %v deleted successfully", transactionID)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": responseMsg,
	})
}
