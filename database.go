// package main

// import (
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// func init() {
// 	var db *gorm.DB
// 	//open a db connection
// 	var err error
// 	db, err = gorm.Open("mysql", "root:***REMOVED***@/godo?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	//Migrate the schema
// 	db.AutoMigrate(&todoModel{})
// }

// type (
// 	// todoModel describes a todoModel type
// 	todoModel struct {
// 		gorm.Model
// 		Title    string `json:"title"`
// 		Competed int    `json:"completed"`
//   }

//   transformedTodo struct {
//     ID        uint   `json:"id"`
//     Title     string `json:"title"`
//     Completed bool   `json:"completed"`
//    }
// )
