package models

import (
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "root:123456a@tcp(localhost:3306)/gingorm?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{}, &Person{}, &IDCard{})
	if err != nil {
		panic("Failed to migrate database!")
		return
	}
	DB = database
}
