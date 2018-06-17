package auth

import (
	"log"
	"net/http"

	"github.com/aflashyrhetoric/payup-api/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type (
	// User ... represents a user of the app
	User struct {
		id       int    `db:"id"`
		Email    string `db:"email"`
		password []byte `db:"password"`
	}
)

func getEmail(c *gin.Context) string {
	return c.PostForm("email")
}

// getPassword returns the password from POST as a []byte
func getPassword(c *gin.Context) []byte {
	return []byte(c.PostForm("password"))
}

// Login ... Checks users and crafts and issues JWT tokens
func Login(c *gin.Context) {
	jwt.New(jwt.GetSigningMethod("HMAC"))
}

func hashSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// Register ... Creates a new user if it doesn't exist
func Register(c *gin.Context) {

	// open connection to database
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Retrieve POST values
	email := getEmail(c)
	password := getPassword(c)
	if err != nil {
		log.Print(err)
	}

	// Hash and salt
	hash := hashSalt(password)

	// Build up model to be saved
	user := User{
		Email:    email,
		password: password,
	}

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO users 
		VALUES(NULL, ?, ?)
	`)
	if err != nil {
		log.Print(err)
	}
	res, err := stmt.Exec(
		user.Email,
		user.password)
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
			"message": "User created successfully.",
			"data":    lastInsertID,
		},
	)
}
