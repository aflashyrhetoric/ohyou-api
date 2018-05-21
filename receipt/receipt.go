package model

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aflashyrhetoric/payup-api/database"
	"github.com/aflashyrhetoric/payup-api/utils"
	"github.com/gin-gonic/gin"
)

type (

	// Receipt ... is a single receipt
	Receipt struct {
		id       int    `db:"id"`
		merchant string `db:"merchant"`
		total    int    `db:"total"`
	}
)

// Get beneficiaries from POST variable
func getMerchant(c *gin.Context) string {
	return c.PostForm("merchant")
}

func getID(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param("id"))
}

func getTotal(c *gin.Context) (int, error) {
	return utils.ConvertDollarsStringToCents(c.PostForm("total"))
}

// CreateReceipt ...router method for creating a new receipt
func CreateReceipt(c *gin.Context) {

	// Initial expense
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Retrieve POST values
	merchant := getMerchant(c)
	total, err := getTotal(c)
	if err != nil {
		log.Print(err)
	}

	// Build up model to be saved
	newReceipt := Receipt{
		merchant: merchant,
		total:    total,
	}

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO receipts 
		VALUES(NULL, ?, ?)
	`)
	if err != nil {
		log.Print(err)
	}
	res, err := stmt.Exec(
		newReceipt.merchant,
		newReceipt.total)
	if err != nil {
		log.Print(err)
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	// Response
	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Receipt created successfully.",
			"data":    lastInsertID,
		},
	)
}

// ShowReceipt ...show a receipt based on its ID
func ShowReceipt(c *gin.Context) {
	var (
		id           int
		merchant     string
		total        int
		responseData Receipt
	)
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	receiptID, err := getID(c)
	if err != nil {
		log.Print(err)
	}
	// Check for invalid ID
	if receiptID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Receipt found!"})
		return
	}
	// Prepare SELECT statement
	stmt, err := db.Prepare(`
		SELECT id, merchant, total
		FROM receipts
		WHERE id=?
	`)
	if err != nil {
		log.Print(err)
	}

	// Run Query
	row := stmt.QueryRow(receiptID)
	if err != nil {
		log.Print(err)
	}

	// Scan values to Go variables
	err = row.Scan(&id, &merchant, &total)
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

	responseData = Receipt{id, merchant, total}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responseData})
}

// UpdateReceipt ... will update a receipt based on its ID
func UpdateReceipt(c *gin.Context) {

	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Retrieve POST update data
	receiptID, _ := getID(c)
	merchant := getMerchant(c)
	total, _ := getTotal(c)

	// Check for invalid ID
	if receiptID == 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No Receipt found!",
			})
		return
	}
	// Prepare SELECT statement
	tx, err := db.Begin()
	stmt, err := tx.Prepare(`
		UPDATE receipts 
		SET merchant=?, total=?
		WHERE id=?;
	`)
	if err != nil {
		log.Print(err)
	}

	// Execute update query for regular expense data
	_, err = stmt.Exec(merchant, total, receiptID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Receipt updated successfully.",
		},
	)
}

// DeleteReceipt ... Deletes a receipt by its ID
func DeleteReceipt(c *gin.Context) {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	receiptID, err := getID(c)
	if err != nil {
		log.Print(err)
	}
	// Check for invalid ID
	if receiptID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Receipt found!"})
		return
	}
	// Prepare SELECT statement
	tx, err := db.Begin()
	stmt, err := tx.Prepare(`
		DELETE FROM expenses
		WHERE id=?
	`)
	if err != nil {
		log.Print(err)
	}

	// Run Query
	_, err = stmt.Exec(receiptID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	responseMsg := fmt.Sprintf("Receipt %v deleted successfully", receiptID)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": responseMsg,
	})
}
