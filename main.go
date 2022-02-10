package main

import (
	"inventory-system/db"
	"inventory-system/models"
	"inventory-system/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database := db.SetupConn()
	database.AutoMigrate(models.User{}, models.Product{})

	user_rt := routes.User{
		Conn: database,
	}

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", MainRoute)

	user_gr := app.Group("/user")
	{
		user_gr.GET("", user_rt.GetAll)
		user_gr.POST("", user_rt.CreateOne)
		user_gr.GET("/:id", user_rt.GetOne)
		user_gr.PUT("/:id", user_rt.UpdateOne)
		user_gr.DELETE("/:id", user_rt.DeleteOne)
	}

	app.Logger.Fatal(app.Start(":8000"))
}

func MainRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Inventory Management System.")
}
