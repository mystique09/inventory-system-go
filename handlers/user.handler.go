package handlers

import (
	"inventory-system/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func GetUserByUsername(conn *gorm.DB, payload *models.ULoginPayload) models.User {
	var result models.User

	conn.Model(&models.User{}).Where("username = ?", payload.Username).Find(&result)
	return result
}

func CreateUser(conn *gorm.DB, payload models.CreateUserDto) error {
	new_user := models.NewUser(&payload)
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(new_user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	new_user.Password = string(hashed_password)

	if err := conn.Create(&new_user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(conn *gorm.DB, uid uuid.UUID, payload models.UpdateUserDto) error {
	var user models.User

	if err := conn.Model(&user).Where("id = ?", uid).Updates(models.User{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
	}).Error; err != nil {
		return err
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
