package handlers

import (
	"inventory-system/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetProducts(conn *gorm.DB) []models.Product {
	var result []models.Product

	conn.Model(models.Product{}).Scan(&result)

	return result
}

func GetProduct(conn *gorm.DB, id uuid.UUID) models.Product {
	var result models.Product

	conn.Model(&models.Product{}).Where("id = ?", id).Find(&result)

	return result
}

func CreateProduct(conn *gorm.DB, payload models.CreateProductDto) error {

	new_product := models.NewProduct(payload)

	if err := conn.Create(&new_product).Error; err != nil {
		return err
	}

	return nil
}

func UpdateProduct(conn *gorm.DB, id uuid.UUID, payload models.UpdateProductDto) error {
	var product models.Product
	if err := conn.Model(&product).Where("id = ?", id).Updates(models.Product{
		Name:  payload.Name,
		Brand: payload.Brand,
		Price: payload.Price,
		Stock: payload.Stock,
	}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteProduct(conn *gorm.DB, id uuid.UUID) error {
	var product models.Product

	if err := conn.Delete(&product, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
