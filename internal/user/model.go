package user

import (
	"time"

	"gorm.io/gorm"
)

// User представляет пользователя в системе
// @Description Модель пользователя с email и паролем
type User struct {
	ID        uint           `json:"id" gorm:"primarykey" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggertype:"string" example:"2023-01-01T00:00:00Z"`
	Email     string         `json:"email" gorm:"uniqueIndex" example:"user@example.com"`
	Password  string         `json:"password" example:"password123"`
}
