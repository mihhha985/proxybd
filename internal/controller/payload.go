package controller

import "test/internal/repository"

// ListResponse представляет ответ со списком пользователей
// @Description Ответ с общим количеством и списком пользователей
type ListResponse struct {
	Total int64             `json:"total" example:"100"`
	Users []repository.User `json:"users"`
}

// CreateUserRequest представляет данные для создания пользователя
// @Description Структура запроса для создания нового пользователя
type CreateUserRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"securepassword123"`
}

// UpdateUserRequest представляет данные для обновления пользователя
// @Description Структура запроса для обновления пользователя
type UpdateUserRequest struct {
	Email    string `json:"email,omitempty" example:"newemail@example.com"`
	Password string `json:"password,omitempty" example:"newsecurepassword123"`
}

type UserResponse struct {
	ID        uint    `json:"id" example:"1"`
	CreatedAt string  `json:"created_at" example:"2024-01-01T12:00:00Z"`
	UpdatedAt string  `json:"updated_at" example:"2024-01-02T12:00:00Z"`
	DeletedAt *string `json:"deleted_at,omitempty" example:"2024-01-03T12:00:00Z"`
	Email     string  `json:"email" example:"user@example.com"`
	Password  string  `json:"password" example:"securepassword123"`
}
