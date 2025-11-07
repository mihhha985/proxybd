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
