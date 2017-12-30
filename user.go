package main

type User struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`

	Purchases []Transaction `gorm:"ForeignKey:Purchaser"`
}
