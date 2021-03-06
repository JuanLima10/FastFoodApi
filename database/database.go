package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "root:123456@tcp(127.0.0.1:3306)/fastfood?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error", err)
	}

	db = database
}

func GetDatabase() *gorm.DB {
	return db
}
