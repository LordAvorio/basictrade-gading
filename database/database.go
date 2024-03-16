package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartDB() (*gorm.DB, error) {

	serverHost := viper.GetString("DB_HOST")
	serverPort := viper.GetString("DB_PORT")
	databaseUsername := viper.GetString("DB_USER")
	databasePassword := viper.GetString("DB_PASS")
	databaseName := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseUsername, databasePassword, serverHost, serverPort, databaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}