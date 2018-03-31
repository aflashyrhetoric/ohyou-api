package model

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aflashyrhetoric/payup-api/database"
	"github.com/aflashyrhetoric/payup-api/utils"
	"github.com/gin-gonic/gin"
)

type (

	// Expense ... is a single purchase
	Expense struct {
		id          int    `db:"id"`
		description string `db:"description"`
		purchaser   int    `db:"purchaser"`
		amount      int    `db:"amount"`
		receipt_id  int    `db:"receipt_id"`
	}
	// TransformedExpense ... is a Expense with additional information
	transformedExpense struct {
		ID            int    `json:"id"`
		Description   string `json:"description"`
		Purchaser     int    `json:"purchaser"`
		Amount        int    `json:"amount"`
		Beneficiaries []int  `json:"beneficiaries"`
		ReceiptID     int    `json:"receipt_id"`
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
func loadExpenseBeneficiaryData(expenseID int) []int {
	var beneficiaryIDs []int
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		SELECT * FROM expenses 
		WHERE expense_id = ?
	`)
	if err != nil {
		log.Print(err)
	}
	rows, err := stmt.Query(expenseID)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	// Scan values to Go variables
	for rows.Next() {
		err := rows.Scan(&expenseID, beneficiaryIDs)
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
	return utils.ConvertDollarsStringToCents(c.PostForm("amount"))
}

func getReceiptID(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param("receipt_id"))
}

// CreateExpense ... Saves an Expense to the DB
func CreateExpense(c *gin.Context) {

	// Initial expense
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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
	receiptID, err := getReceiptID(c)
	if err != nil {
		log.Print(err)
	}

	// Build up model to be saved
	newExpense := transformedExpense{
		Description:   description,
		Purchaser:     int(purchaser),
		Amount:        amount,
		Beneficiaries: beneficiaries,
		ReceiptID:     receiptID,
	}

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO expenses 
		VALUES(NULL, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Print(err)
	}
	res, err := stmt.Exec(
		newExpense.Description,
		newExpense.Amount,
		newExpense.Purchaser,
		newExpense.ReceiptID)
	if err != nil {
		log.Print(err)
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	// Save beneficiaries data

	// Save a expense (to the junction table) for each beneficiary
	for _, beneficiaryID := range newExpense.Beneficiaries {
		tx, err = db.Begin()
		if err != nil {
			log.Print(err)
		}
		stmt, err = tx.Prepare(`
			INSERT INTO expenses_beneficiaries 
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
			"message": "Expense created successfully.",
			"data":    lastInsertID,
		},
	)
}

//ListExpenses ...Lists all current expenses
func ListExpenses(c *gin.Context) {
	var (
		ID            int
		Description   string
		Purchaser     int
		Amount        int
		BeneficiaryID int
		Beneficiaries []int
		ReceiptID     int
		responseData  []transformedExpense
	)
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare SELECT statement
	stmt, err := db.Prepare("SELECT * FROM expenses")
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
		err := rows.Scan(&ID, &Description, &Purchaser, &Amount, &ReceiptID)
		if err != nil {
			log.Print(err)
		}

		// Run second query to retrieve expenses_beneficiaries data
		stmt, err = db.Prepare(`
			SELECT beneficiary_id 
			FROM expenses_beneficiaries 
			WHERE expense_id = ?
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
		responseData = append(responseData, transformedExpense{ID, Description, Purchaser, Amount, Beneficiaries, ReceiptID})
		Beneficiaries = nil
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responseData})
}

// ShowExpense shows a single expense
func ShowExpense(c *gin.Context) {
	var (
		ID            int
		Description   string
		Purchaser     int
		Amount        int
		BeneficiaryID int
		Beneficiaries []int
		ReceiptID     int
		responseData  transformedExpense
	)
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ExpenseID, err := getID(c)
	if err != nil {
		log.Print(err)
	}
	// Check for invalid ID
	if ExpenseID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Expense found!"})
		return
	}
	// Prepare SELECT statement
	stmt, err := db.Prepare(`
		SELECT id, description, purchaser, amount, receipt_id
		FROM expenses
		WHERE id=?
	`)
	if err != nil {
		log.Print(err)
	}

	// Run Query
	row := stmt.QueryRow(ExpenseID)
	if err != nil {
		log.Print(err)
	}

	// Scan values to Go variables
	err = row.Scan(&ID, &Description, &Purchaser, &Amount, &ReceiptID)
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

	// Retrieve expenses_beneficiaries data
	stmt, err = db.Prepare(`
		SELECT beneficiary_id 
		FROM expenses_beneficiaries 
		WHERE expense_id = ?
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

	responseData = transformedExpense{ID, Description, Purchaser, Amount, Beneficiaries, ReceiptID}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responseData})
}

// UpdateExpense updates a single expense in db with new fields
func UpdateExpense(c *gin.Context) {

	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Retrieve POST update data
	expenseID, _ := getID(c)
	description := getDescription(c)
	purchaser, _ := getPurchaser(c)
	amount, _ := getAmount(c)
	receiptID, _ := getReceiptID(c)

	// Check for invalid ID
	if expenseID == 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No Expense found!",
			})
		return
	}
	// Prepare SELECT statement
	tx, err := db.Begin()
	stmt, err := tx.Prepare(`
		UPDATE expenses 
		SET description=?, purchaser=?, amount=?, receipt=?
		WHERE id=?;
	`)
	if err != nil {
		log.Print(err)
	}

	// Execute update query for regular expense data
	_, err = stmt.Exec(description, purchaser, amount, receiptID, expenseID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	// Update expenses_beneficiaries data
	beneficiaries := getBeneficiaries(c)

	// First, delete all old associated beneficiaries
	tx, err = db.Begin()
	stmt, err = tx.Prepare(`
		DELETE FROM expenses_beneficiaries 
		WHERE expense_id=?;
	`)
	if err != nil {
		log.Print(err)
	}
	// Execute update query for beneficiaries
	_, err = stmt.Exec(expenseID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	// Second, insert new values
	for _, beneficiaryID := range beneficiaries {
		tx, err = db.Begin()
		stmt, err = tx.Prepare(`
			INSERT INTO expenses_beneficiaries 
			VALUES (?, ?);
		`)
		if err != nil {
			log.Print(err)
		}

		// Execute update query for beneficiaries
		_, err = stmt.Exec(expenseID, beneficiaryID)
		if err != nil {
			log.Print(err)
		}
		tx.Commit()
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Expense updated successfully.",
		},
	)
}

func DeleteExpense(c *gin.Context) {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ExpenseID, err := getID(c)
	if err != nil {
		log.Print(err)
	}
	// Check for invalid ID
	if ExpenseID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Expense found!"})
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
	_, err = stmt.Exec(ExpenseID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	// Delete beneficiary data
	tx, err = db.Begin()
	stmt, err = tx.Prepare(`
		DELETE FROM expenses_beneficiaries
		WHERE expense_id=?
	`)
	if err != nil {
		log.Print(err)
	}

	// Run Query
	_, err = stmt.Exec(ExpenseID)
	if err != nil {
		log.Print(err)
	}
	tx.Commit()

	responseMsg := fmt.Sprintf("Expense %v deleted successfully", ExpenseID)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": responseMsg,
	})
}
