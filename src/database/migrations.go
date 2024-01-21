package database

import (
	"go-marketplace/src/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(model.User{}, model.Address{})
	return err
}
