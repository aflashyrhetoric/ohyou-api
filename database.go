package main

import (
	"log"

	db "upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

func main() {

	// Configure connection
	var settings = mysql.ConnectionURL{
		User:     "root",
		Password: "password",
		Host:     "localhost",
		Database: "ohyou_api",
	}

	// Open connection
	// sess, err := mysql.Open(settings)
	sess, err := db.Open(mysql.Adapter, settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer sess.Close()
}
