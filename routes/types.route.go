package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Route interface {
		// get all from database
		InitDb(*gorm.DB)
		GetAll(echo.Context) []interface{}
		GetOne(echo.Context) interface{}
		CreateOne(echo.Context) error
		UpdateOne(echo.Context) error
		DeleteOne(echo.Context) error
	}
)

func CreateRoute(conn *gorm.DB, rt Route) interface{} {
	return rt
}
