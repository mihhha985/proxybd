package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"test/pkg/response"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	userRepository IUserRepository
}

func NewUserController(userRepository IUserRepository) *UserController {
	return &UserController{
		userRepository: userRepository,
	}
}

// GetOne получает пользователя по ID
// @Summary Получить пользователя по ID
// @Description Возвращает информацию о пользователе по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} user.User "Пользователь найден"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Router /users/{id} [get]
func (uc *UserController) GetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := uc.userRepository.GetByID(r.Context(), id)
	if err != nil {
		response.ErrorNotFound(w, err)
		return
	}
	response.OutputJSON(w, user)
}

// GetAll получает список всех пользователей с пагинацией
// @Summary Получить список пользователей
// @Description Возвращает список всех пользователей с поддержкой пагинации
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int true "Количество записей" example(10)
// @Param offset query int true "Смещение" example(0)
// @Success 200 {object} user.ListResponse "Список пользователей"
// @Failure 400 {object} map[string]string "Некорректные параметры запроса"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Header 200 {string} X-Total-Count "Общее количество пользователей"
// @Router /users [get]
func (uc *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		response.ErrorBadRequest(w, errors.New("неверные параметры запроса"))
		return
	}
	offset := r.URL.Query().Get("offset")
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		response.ErrorBadRequest(w, errors.New("неверные параметры запроса"))
		return
	}

	conditions := Conditions{
		Limit:  limitInt,
		Offset: offsetInt,
	}
	count := uc.userRepository.Count()
	w.Header().Set("X-Total-Count", strconv.FormatInt(count, 10))
	users, err := uc.userRepository.List(r.Context(), conditions)
	if err != nil {
		response.ErrorInternal(w, err)
	}
	resp := ListResponse{
		Total: count,
		Users: users,
	}
	response.OutputJSON(w, resp)
}

// Create создает нового пользователя
// @Summary Создать нового пользователя
// @Description Создает нового пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param user body user.CreateUserRequest true "Данные пользователя"
// @Success 200 {object} user.User "Пользователь создан"
// @Failure 400 {object} map[string]string "Некорректные данные"
// @Router /users [post]
func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var data CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.ErrorBadRequest(w, err)
		return
	}
	user := User{
		Email:    data.Email,
		Password: data.Password,
	}
	err = uc.userRepository.Create(r.Context(), user)
	if err != nil {
		response.ErrorBadRequest(w, err)
		return
	}
	response.OutputJSON(w, user)
}

// Update обновляет данные пользователя
// @Summary Обновить пользователя
// @Description Обновляет данные существующего пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Param user body user.UpdateUserRequest true "Обновленные данные пользователя"
// @Success 200 {object} user.User "Пользователь обновлен"
// @Failure 400 {object} map[string]string "Некорректные данные"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Router /users/{id} [put]
func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Преобразуем ID в uint
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.ErrorBadRequest(w, errors.New("некорректный ID"))
		return
	}

	// Проверяем, существует ли пользователь
	existingUser, err := uc.userRepository.GetByID(r.Context(), id)
	if err != nil {
		response.ErrorNotFound(w, errors.New("пользователь не найден"))
		return
	}

	var data UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.ErrorBadRequest(w, err)
		return
	}

	// Обновляем только переданные поля
	existingUser.ID = uint(idInt)
	if data.Email != "" {
		existingUser.Email = data.Email
	}
	if data.Password != "" {
		existingUser.Password = data.Password
	}

	err = uc.userRepository.Update(r.Context(), existingUser)
	if err != nil {
		response.ErrorBadRequest(w, err)
		return
	}
	response.OutputJSON(w, existingUser)
}

// Delete удаляет пользователя по ID
// @Summary Удалить пользователя
// @Description Удаляет пользователя из системы по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} map[string]string "Пользователь удален"
// @Failure 400 {object} map[string]string "Ошибка при удалении"
// @Router /users/{id} [delete]
func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := uc.userRepository.Delete(r.Context(), id)
	if err != nil {
		response.ErrorBadRequest(w, err)
		return
	}
	response.OutputJSON(w, map[string]string{"result": "success"})
}
