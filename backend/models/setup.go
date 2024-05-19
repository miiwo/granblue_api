package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm"
	"fmt"
	"os"
)
var DB *gorm.DB

func ConnectToMariaDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	//dsn := "root:password@tcp(localhost:3306)/weapons_and_charas?charset=utf8&parseTime=True&loc=Local"
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