package main

import (
	"go-marketplace/src/api"
	"go-marketplace/src/database"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	app := echo.New()
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

	err = database.RunSeeders(db)
	if err != nil {
		log.Fatal("failed to run seeders", err)
	}

	api.SetupRoutes(app, db)

}
