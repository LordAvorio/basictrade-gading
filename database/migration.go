package database

import "gorm.io/gorm"

func RunMigration(db *gorm.DB) error {

	err := db.AutoMigrate()

	if err != nil {
		return err
	}

	return nil
}