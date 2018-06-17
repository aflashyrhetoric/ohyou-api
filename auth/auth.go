package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type (
	// User ... represents a user of the app
	User struct {
		id       int    `db:"id"`
		email    string `db:"email"`
		password string `db:"password"`
	}
)

func getEmail(c *gin.Context) string {
	return c.PostForm("email")
}

func getPassword(c *gin.Context) string {
	return c.PostForm("password")
}

// Register ... Creates a new user if it doesn't exist
func Register(c *gin.Context) {
	// jwt.
}

// Login ... Checks users and crafts and issues JWT tokens
func Login(c *gin.Context) {
	jwt.New()
}
