package main

import "github.com/jinzhu/gorm"

type user struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
