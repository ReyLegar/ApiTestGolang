package rest

import (
	"ApiTest/internal/models"
	"ApiTest/internal/services"
	"encoding/json"
	"io"
	"net/http"
)

type UserHandlers struct {
	userService services.UserService
}

func NewUserHandlers(userService services.UserService) *UserHandlers {
	return &UserHandlers{userService: userService}
}

func (h *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Не тот метод", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Ошибка при разборе JSON", http.StatusBadRequest)
		return
	}

	id, err := h.userService.CreateUser(&user)
	if err != nil {
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]int64{"Успех": id}
	json.NewEncoder(w).Encode(response)
}
