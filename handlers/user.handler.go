package handlers

import (
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
	conn.Model(&models.User{}).
		Where("id = ?", id).
		Find(&result)

	return result
}

func CreateUser(conn *gorm.DB, payload models.CreateUserDto) error {
	new_user := models.NewUser(&payload)
	if err := conn.Create(&new_user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(conn *gorm.DB, uid uuid.UUID, payload models.UpdateUserDto, field string) error {
	var user models.User

	switch field {
	case "username":
		err := conn.Model(&user).Where("id = ?", uid).Update("username", payload.Username).Error
		if err != nil {
			return err
		}
	case "password":
		err := conn.Model(&user).Where("id = ?", uid).Update("password", payload.Password).Error

		if err != nil {
			return err
		}

	case "email":
		err := conn.Model(&user).Where("id = ?", uid).Update("email", payload.Email).Error

		if err != nil {
			return err
		}

	}

	return nil
}

func DeleteUser(conn *gorm.DB, uid uuid.UUID) error {
	var deleted_user models.User

	if err := conn.Delete(&deleted_user, "id = ?", uid).Error; err != nil {
		return err
	}
	return nil
}
