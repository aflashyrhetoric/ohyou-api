package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		ID            int    `json:id`
		Description   string `json:"description"`
		Purchaser     int    `json:"purchaser"`
		Amount        int    `json:"amount"`
		Beneficiaries []User `json:"beneficiaries"`
	}
)

func getBeneficiaries(c *gin.Context) []User {
	u := User{}
	u.Name = "bob"
	u.Email = "aflashyrhetoric@gmail.com"
	u.Password = "password"

	var beneficiaries []User

	beneficiaries = append(beneficiaries, u)
	beneficiaries = append(beneficiaries, u)
	beneficiaries = append(beneficiaries, u)
	beneficiaries = append(beneficiaries, u)

	return beneficiaries
}

func getID(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param("id"))
}

func getDescription(c *gin.Context) string {
	return c.PostForm("description")
}

func getPurchaser(c *gin.Context) (int64, error) {
	return strconv.ParseInt(c.PostForm("purchaser"), 10, 0)
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
		log.Fatal(err)
	}
	amount, err := getAmount(c)
	if err != nil {
		log.Fatal(err)
	}
	// Build up model to be saved
	newTransaction := transformedTransaction{
		Description: description,
		Purchaser:   int(purchaser),
		Amount:      amount,
	}

	// The ID of the transaction AFTER it is saved to the transactions table
	var newTransactionID int64

	// Save initial transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO transactions VALUES(NULL, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(
		newTransaction.Description,
		newTransaction.Amount,
		newTransaction.Purchaser)
	if err != nil {
		log.Fatal(err)
	} else {
		newTransactionID, err = res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	// Save beneficiaries data

	// Re-assign tx and err variables from earlier

	// TODO: setup a loop for each beneficiary, run a new transaction to save to the database
	tx, err = db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err = tx.Prepare("INSERT INTO transactions_beneficiaries VALUES(NULL, ?, ?)")
	res, err = stmt.Exec(
		newTransaction.Description,
		newTransaction.Amount,
		newTransaction.Purchaser)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	// Response
	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Transaction created successfully.",
			"data":    res.LastInsertId,
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
	// Prepare SELECT statement
	stmt, err := db.Prepare("SELECT * FROM transactions")
	if err != nil {
		log.Fatal(err)
	}
	// Run Query
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// Scan values to Go variables
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

func showTransaction(c *gin.Context) {
	var (
		ID           int
		Description  string
		Purchaser    int
		Amount       int
		responseData transformedTransaction
	)
	transactionID, err := getID(c)
	if err != nil {
		log.Fatal(err)
	}
	// Check for invalid ID
	if transactionID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}
	// Prepare SELECT statement
	stmt, err := db.Prepare("SELECT id, description, purchaser, amount FROM transactions WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	// Run Query
	row := stmt.QueryRow(transactionID)
	if err != nil {
		log.Fatal(err)
	}

	// Scan values to Go variables
	err = row.Scan(&ID, &Description, &Purchaser, &Amount)
	if err != nil {
		log.Fatal(err)
	}
	responseData = transformedTransaction{ID, Description, Purchaser, Amount, nil}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responseData})
}

func updateTransaction(c *gin.Context) {
	transactionID, err := getID(c)
	description := getDescription(c)
	purchaser, err := getPurchaser(c)
	amount, err := getAmount(c)
	if err != nil {
		log.Fatal(err)
	}
	// Check for invalid ID
	if transactionID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
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
		log.Fatal(err)
	}

	// Run Query
	_, err = stmt.Exec(description, purchaser, amount, transactionID)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	// Run Query
	_, err = stmt.Exec(transactionID)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	responseMsg := fmt.Sprintf("Transaction %v deleted successfully", transactionID)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": responseMsg,
	})
}
