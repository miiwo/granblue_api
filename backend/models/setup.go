package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm"
	"fmt"
)
var DB *gorm.DB

func ConnectToMariaDB() (*gorm.DB, error) {
	dsn := "root:password@tcp(localhost:3306)/weapons_and_charas?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Could not connect to DB :(")
		return nil, err
	}

	fmt.Println("Succesfully connected to DB!")
	DB = db

	return db, nil
}