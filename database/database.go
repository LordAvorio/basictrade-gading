package database

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartDB() (*gorm.DB, error) {

	serverHost := os.Getenv("DB_HOST")
	serverPort := os.Getenv("DB_PORT")
	databaseUsername := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASS")
	databaseName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseUsername, databasePassword, serverHost, serverPort, databaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}