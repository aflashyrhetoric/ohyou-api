package database

import "github.com/aflashyrhetoric/ohyou-api/model"

var DatabaseInfo model.Database

func NewDB(){
	// Initialize db
	var err error
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/ohyou_api")
	if err != nil {
		log.Fatal(err)
  }
  
  return db
}

func InitDatabase() {
  db :=
}