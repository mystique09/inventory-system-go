package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	// The product model.
	Product struct {
		gorm.Model
		Name     string    `json:"product_name"`
		Image    string    `json:"product_image"`
		Brand    string    `json:"product_brand"`
		Price    uint32    `json:"product_price"`
		Category string    `json:"product_category"`
		Owner    uuid.UUID `json:"product_owner"`
		Stock    uint32    `json:"product_stock"`
	}

	// The product DTO, use for request binding
	// when creating new Product.
	CreateProductDto struct {
		Name     string
		Image    string
		Brand    string
		Price    uint32
		Category string
		Owner    uuid.UUID
		Stock    uint32
	}

	// Use for updating existing Product.
	UpdateProductDto struct {
		Name  string
		Brand string
		Price string
		Stock uint32
	}
)

func (*Product) New(payload CreateProductDto) Product {
	default_image := "https://default-image-url.com"

	if payload.Image == "" {
		payload.Image = default_image
	}

	new_product := Product{
		Name:     payload.Name,
		Image:    payload.Image,
		Brand:    payload.Brand,
		Price:    payload.Price,
		Category: payload.Category,
		Owner:    payload.Owner,
		Stock:    payload.Stock,
	}

	return new_product
}
