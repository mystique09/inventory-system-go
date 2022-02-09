package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string
	Password string
	Email    string
	Role     string
}

type CreateUserDto struct {
	Username string
	Password string
	Email    string
	Role     string
}

type UpdateUserDto struct {
	Username string
	Password string
	Email    string
	Role     string
}
