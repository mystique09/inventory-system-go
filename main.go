package main

import (
	"fmt"

	"inventory-system/pkg/db"
	"inventory-system/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conn := db.SetupConn()
	conn.AutoMigrate(models.User{}, models.Product{})

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	user := app.Group("/user", middleware.JWT(func(key string) {
		fmt.Println(key)
	}))
	user.GET("/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})

	app.Logger.Fatal(app.Start(":8000"))
}
