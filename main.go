package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/transactions")
	{
		v1.GET("/", listTransaction)
		v1.POST("/", createTransaction)
		// v1.GET("/:id", showTransaction)
		// v1.PUT("/:id", updateTransaction)
		// v1.DELETE("/:id", deleteTransaction)
	}
	router.Run()
}

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:password@/ohyou_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&transaction{})
}
