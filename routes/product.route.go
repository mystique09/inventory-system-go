package routes

import (
	"inventory-system/handlers"
	"inventory-system/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Product struct {
		Conn *gorm.DB
	}
)

func (rt *Product) InitDb(conn *gorm.DB) {
	rt.Conn = conn
}

func (rt *Product) GetAll(c echo.Context) error {
	var products []models.Product = handlers.GetProducts(rt.Conn)

	return c.JSON(http.StatusOK, products)
}

func (rt *Product) CreateOne(c echo.Context) error {

	payload := new(models.CreateProductDto)

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if payload.Owner == uuid.Nil {
		return c.JSON(http.StatusBadRequest, "Missing owner field.")
	}

	if err := handlers.CreateProduct(rt.Conn, *payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "One product created.")
}

func (rt *Product) GetOne(c echo.Context) error {
	uid, uid_error := uuid.Parse(c.Param("id"))

	if uid_error != nil {
		return c.JSON(http.StatusBadRequest, uid_error)
	}

	var product models.Product = handlers.GetProduct(rt.Conn, uid)
	return c.JSON(http.StatusOK, product)
}

func (rt *Product) UpdateOne(c echo.Context) error {
	uid, uid_err := uuid.Parse(c.Param("id"))
	payload := new(models.UpdateProductDto)

	if uid_err != nil {
		return c.JSON(http.StatusBadRequest, uid_err)
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := handlers.UpdateProduct(rt.Conn, uid, *payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "Updated one product.")
}

func (rt *Product) DeleteOne(c echo.Context) error {

	uid, uid_err := uuid.Parse(c.Param("id"))

	if uid_err != nil {
		return c.JSON(http.StatusBadRequest, uid_err)
	}

	if err := handlers.DeleteProduct(rt.Conn, uid); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "Deleted one product.")
}
