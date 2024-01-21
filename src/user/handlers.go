package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	r UserRepository
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{
		r: repo,
	}
}

func (u *UserHandler) ListUsers(c echo.Context) error {
	users, err := u.r.FindAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (u *UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "could not parse user id")
	}

	user, err := u.r.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"message": "did not find user with specified id", "success": false})
	}

	return c.JSON(http.StatusOK, user)
}
