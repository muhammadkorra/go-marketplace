package database

import (
	"go-marketplace/src/address"
	"go-marketplace/src/order"
	"go-marketplace/src/product"
	"go-marketplace/src/user"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(user.User{}, address.Address{}, product.Product{}, order.Order{}, order.OrderItem{})
	return err
}

func RunSeeders(db *gorm.DB) error {
	productSeeder := product.NewProductsSeeder(db)
	return productSeeder.GenerateProducts()
}
