package database

import (
	"go-marketplace/src/address"
	"go-marketplace/src/user"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(user.User{}, address.Address{})
	return err
}
