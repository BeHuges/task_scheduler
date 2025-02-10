package handler

import (
	"encoding/json"
	constants "go_final_project/internal/constants"
	"go_final_project/internal/error"
	nextdate "go_final_project/internal/next_date"
	"go_final_project/internal/task"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	var t task.Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "JSON Deserialization Error", http.StatusBadRequest)
		return
	}

	//Проверяем id
	if t.ID == "" {
		error.JsonResponse(w, "Id задачи не указан")
		return
	}

	//Проверяем title
	if t.Title == "" {
		error.JsonResponse(w, "Не указан заголовок задачи")
		return
	}

	//Проверяем поле date
	now := time.Now()
	if t.Date == "" {
		t.Date = now.Format(constants.DateFormat)
	} else {
		parseDate, err := time.Parse(constants.DateFormat, t.Date)
		if err != nil {
			error.JsonResponse(w, "Дата указана в неверном формате")
			return
		}
		if parseDate.Before(now) && t.Repeat != "" {
			nextDate, err := nextdate.NextDate(now, t.Date, t.Repeat)
			if err != nil {
				error.JsonResponse(w, "Неверный формат правила повторения")
				return
			}
			t.Date = nextDate
		}
	}

	//Проверяем правило повторения
	if t.Repeat != "" {
		_, err := nextdate.NextDate(now, t.Date, t.Repeat)
		if err != nil {
			error.JsonResponse(w, "Неверный формат правила повторения")
			return
		}
	}

	//Конвертируем id в int
	idInt, err := strconv.Atoi(t.ID)
	if err != nil {
		error.JsonResponse(w, "Ошибка конвертации id в Int")
		return
	}

	//Update запрос к db
	res, err := h.repo.UpdateTask(t.Date, t.Title, t.Comment, t.Repeat, idInt)
	if err != nil || res == 0 {
		error.JsonResponse(w, "Ошибка обновления задачи")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{}"))
}
