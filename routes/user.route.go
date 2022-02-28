package routes

import (
	"fmt"
	"inventory-system/handlers"
	"inventory-system/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	Conn *gorm.DB
}

func (rt *User) InitDb(conn *gorm.DB) {
	rt.Conn = conn
}

func (rt *User) GetAll(c echo.Context) error {
	var results []models.UserResponse = handlers.GetUsers(rt.Conn)

	return c.JSON(http.StatusOK, results)
}

func (rt *User) GetOne(c echo.Context) error {
	uid, uuidparse_err := uuid.Parse(c.Param("id"))

	if uuidparse_err != nil {
		return c.JSON(http.StatusBadRequest, uuidparse_err.Error())
	}

	var result models.UserResponse = handlers.GetUser(rt.Conn, uid)

	if result.Username == "" {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("User with id %s doesn't exist.", uid))
	}

	return c.JSON(http.StatusOK, result)
}

func (rt *User) CreateOne(c echo.Context) error {
	payload := new(models.CreateUserDto)

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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
	payload := new(models.UpdateUserDto)
	uid, uuidparse_err := uuid.Parse(c.Param("id"))

	if uuidparse_err != nil {
		return c.JSON(http.StatusBadRequest, uuidparse_err.Error())
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := handlers.UpdateUser(rt.Conn, uid, *payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "User updated.")
}

func (rt *User) DeleteOne(c echo.Context) error {
	uid, uuidparse_err := uuid.Parse(c.Param("id"))
	if uuidparse_err != nil {
		return c.JSON(http.StatusBadRequest, uuidparse_err.Error())
	}

	if err := handlers.DeleteUser(rt.Conn, uid); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "User deleted.")
}
