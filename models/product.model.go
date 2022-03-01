package models

import (
	"github.com/google/uuid"
)

type (
	// The product model.
	Product struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Image    string    `json:"image"`
		Brand    string    `json:"brand"`
		Price    uint32    `json:"price"`
		Category string    `json:"category"`
		Owner    uuid.UUID `json:"owner"`
		Stock    uint32    `json:"stock"`
	}

	// The product DTO, use for request binding
	// when creating new Product.
	CreateProductDto struct {
		Name     string    `json:"name"`
		Image    string    `json:"image"`
		Brand    string    `json:"brand"`
		Price    uint32    `json:"price"`
		Category string    `json:"category"`
		Owner    uuid.UUID `json:"owner"`
		Stock    uint32    `json:"stock"`
	}

	// Use for updating existing Product.
	UpdateProductDto struct {
		Name  string `json:"name"`
		Brand string `json:"brand"`
		Price uint32 `json:"price"`
		Stock uint32 `json:"stock"`
	}
)

func NewProduct(payload CreateProductDto) Product {
	default_image := "https://default-image-url.com"

	if payload.Image == "" {
		payload.Image = default_image
	}

	new_product := Product{
		ID:       uuid.New(),
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
