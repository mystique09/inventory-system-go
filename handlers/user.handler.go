package handlers

import (
	"errors"
	"inventory-system/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUsers(conn *gorm.DB) []models.UserResponse {
	var results []models.UserResponse
	conn.Model(models.User{}).Scan(&results)

	return results
}

func GetUser(conn *gorm.DB, id uuid.UUID) models.UserResponse {
	var result models.UserResponse
	conn.Model(&models.User{}).Where("id = ?", id).Find(&result)

	return result
}

func CreateUser(conn *gorm.DB, payload models.CreateUserDto) error {
	new_user := models.NewUser(&payload)
	if err := conn.Create(&new_user).Error; err != nil {
		return errors.New(err.Error())
	}
	return nil
}
