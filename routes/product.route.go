package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Product struct {
		Conn *gorm.DB
	}
)

func (rt *Product) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetAll product.")
}

func (rt *Product) CreateOne(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateOne product.")
}

func (rt *Product) GetOne(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetOne product.")
}

func (rt *Product) UpdateOne(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateOne product.")
}

func (rt *Product) DeleteOne(c echo.Context) error {
	return c.JSON(http.StatusOK, "DeleteOne product.")
}
