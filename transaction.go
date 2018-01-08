package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	// Transaction ... is a single purchase
	Transaction struct {
		id          int
		description string
		purchaser   int
		amount      int
	}
	// TransformedTransaction ... is a Transaction with additional information
	transformedTransaction struct {
		ID            int    `json:id`
		Description   string `json:"description"`
		Purchaser     int    `json:"purchaser"`
		Amount        int    `json:"amount"`
		Beneficiaries []User `json:"beneficiaries"`
	}
)

func getDescription(c *gin.Context) string {
	return c.PostForm("description")
}

func getPurchaser(c *gin.Context) (int64, error) {
	return strconv.ParseInt(c.PostForm("purchaser"), 10, 0)
}

func getAmount(c *gin.Context) (int, error) {
	return ConvertDollarsStringToCents(c.PostForm("amount"))
}

func getID(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param("id"))
}

func createTransaction(c *gin.Context) {
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
	// Save
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO Transactions VALUES(NULL, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(
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
	stmt, err := db.Prepare("SELECT * FROM Transactions")
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
	TransactionID, err := getID(c)
	if err != nil {
		log.Fatal(err)
	}
	// Check for invalid ID
	if TransactionID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}
	// Prepare SELECT statement
	stmt, err := db.Prepare("SELECT id, description, purchaser, amount FROM Transactions WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	// Run Query
	row := stmt.QueryRow(TransactionID)
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
	TransactionID, err := getID(c)
	description := getDescription(c)
	purchaser, err := getPurchaser(c)
	amount, err := getAmount(c)
	if err != nil {
		log.Fatal(err)
	}
	// Check for invalid ID
	if TransactionID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}
	// Prepare SELECT statement
	tx, err := db.Begin()
	stmt, err := tx.Prepare(`
		UPDATE Transactions 
		SET description=?, purchaser=?, amount=?
		WHERE id=?;
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Run Query
	_, err = stmt.Exec(description, purchaser, amount, TransactionID)
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
	TransactionID, err := getID(c)
	if err != nil {
		log.Fatal(err)
	}
	// Check for invalid ID
	if TransactionID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
		return
	}
	// Prepare SELECT statement
	tx, err := db.Begin()
	stmt, err := tx.Prepare(`
		DELETE FROM Transactions
		WHERE id=?
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Run Query
	_, err = stmt.Exec(TransactionID)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	responseMsg := fmt.Sprintf("Transaction %v deleted successfully", TransactionID)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": responseMsg,
	})
}
