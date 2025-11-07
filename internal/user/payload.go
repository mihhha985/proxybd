package user

// ListResponse представляет ответ со списком пользователей
// @Description Ответ с общим количеством и списком пользователей
type ListResponse struct {
	Total int64  `json:"total" example:"100"`
	Users []User `json:"users"`
}

// Conditions представляет параметры для пагинации
// @Description Условия для получения списка с пагинацией
type Conditions struct {
	Limit  int `json:"limit" example:"10"`
	Offset int `json:"offset" example:"0"`
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
