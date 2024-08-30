package rest

import (
	"ApiTest/internal/models"
	"ApiTest/internal/services"
	"encoding/json"
	"io"
	"net/http"
)

type WorkHandlers struct {
	workService services.WorkService
}

func NewWorkHandlers(workService services.WorkService) *WorkHandlers {
	return &WorkHandlers{workService: workService}
}

func (h *WorkHandlers) HandleCreateWork(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Не тот метод", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}

	var work models.Work
	err = json.Unmarshal(body, &work)
	if err != nil {
		http.Error(w, "Ошибка при разборе JSON", http.StatusBadRequest)
		return
	}

	id, err := h.workService.CreateWork(&work)
	if err != nil {
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]int64{"Успех": id}
	json.NewEncoder(w).Encode(response)
}
