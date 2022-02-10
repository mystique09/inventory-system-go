package routes

import (
	"inventory-system/handlers"
	"inventory-system/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	Conn *gorm.DB
}

func (rt *User) GetAll(c echo.Context) error {
	var results []models.UserResponse = handlers.GetUsers(rt.Conn)

	return c.JSON(http.StatusOK, results)
}

func (rt *User) GetOne(c echo.Context) error {
	return c.String(http.StatusOK, "Get user by id.")
}

func (rt *User) CreateOne(c echo.Context) error {
	payload := new(models.CreateUserDto)

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if payload.Username == "" || payload.Email == "" || payload.Password == "" {
		return c.JSON(http.StatusBadRequest, "Missing fields.")
	}

	var hasUser models.User

	rt.Conn.Model(&models.User{}).Where("username = ?", payload.Username).Find(&hasUser)

	if hasUser.Username != "" {
		return c.JSON(http.StatusBadRequest, "User already exist.")
	}

	if err := handlers.CreateUser(rt.Conn, *payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, "New user created.")
}

func (rt *User) UpdateOne(c echo.Context) error {
	return c.String(http.StatusOK, "Update one")
}

func (rt *User) DeleteOne(c echo.Context) error {
	return c.String(http.StatusOK, "Delete one.")
}
