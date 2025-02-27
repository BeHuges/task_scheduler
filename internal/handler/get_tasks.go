package handler

import (
	"encoding/json"
	"net/http"

	"task_scheduler/internal/error"
	"task_scheduler/internal/task"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	//Проверяем кол-во задач в db
	var count int
	err := h.repo.Count().Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks := []task.Task{}

	//Получаем список задач при их наличии
	if count > 0 {
		tasks, err = h.repo.GetTasks()
		if err != nil {
			error.JsonResponse(w, "Ошибка получения списка задач")
			return
		}
	}

	//Формируем ответ
	response := map[string]interface{}{
		"tasks": tasks,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
