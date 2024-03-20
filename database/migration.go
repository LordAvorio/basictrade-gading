package database

import (
	"basictrade-gading/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {

	err := db.AutoMigrate(
		&models.Admin{},
	)

	if err != nil {
		return err
	}

	return nil
}