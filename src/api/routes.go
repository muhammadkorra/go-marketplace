package api

import (
	"go-marketplace/src/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(app *echo.Echo, db *gorm.DB) {
	userRepo := user.NewUserDBRepository(db)
	userHandler := user.NewUserHandler(userRepo)

	app.GET("/users", userHandler.ListUsers)
	app.GET("/users/:id", userHandler.GetUserByID)

	app.Logger.Fatal(app.Start(":3000"))
}
