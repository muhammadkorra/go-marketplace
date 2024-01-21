package main

import (
	"go-marketplace/src/database"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=password port=5432 sslmode=disable TimeZone=Asia/Dubai"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("could not connect to db")
	}

	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal("Failed to run migrations")
	}
}
