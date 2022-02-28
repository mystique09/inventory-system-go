package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type (
	User struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Email     string    `json:"email"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CreateUserDto struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	UpdateUserDto struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	UserResponse struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

type (
	Response struct {
		Message string
		Success bool
		Data    interface{}
	}
)

func NewUser(payload *CreateUserDto) User {
	return User{
		ID:        uuid.New(),
		Username:  payload.Username,
		Password:  payload.Password,
		Email:     payload.Email,
		Role:      "Normal",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (payload *User) CheckPassword(given_password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(payload.Password), []byte(given_password)); err != nil {
		return false
	}
	return true
}

func (payload *User) HashPassword() {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	*&payload.Password = string(hashedpassword)
}
