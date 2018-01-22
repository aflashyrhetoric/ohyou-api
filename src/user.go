package main

// User ... represents a single user
type User struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
