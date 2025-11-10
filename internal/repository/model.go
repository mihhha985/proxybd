package repository

import (
	"gorm.io/gorm"
)

// User представляет пользователя в системе
// @Description Модель пользователя с email и паролем
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}
