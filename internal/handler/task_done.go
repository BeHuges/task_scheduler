package handler

import (
	//"database/sql"
	"go_final_project/internal/error"
	nextdate "go_final_project/internal/next_date"

	//"go_final_project/internal/task"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) TaskDone(w http.ResponseWriter, r *http.Request) {

	//Проверяем id
	id := r.URL.Query().Get("id")
	if id == "" {
		error.JsonResponse(w, "Отсутсвует указанный id")
		return
	}

	//var t task.Task

	t, err := h.repo.GetTaskByID(id)
	if err != nil {
		error.JsonResponse(w, "Задача не найдена")
		return
	}

	//Получаем значение полей по id
	//row := h.repo.GetTaskByID(id)
	//err := row.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//	} else {
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//	}
	//	return
	//}

	//Конвертируем id в int
	idInt, err := strconv.Atoi(t.ID)
	if err != nil {
		error.JsonResponse(w, "Ошибка конвертации id в int")
		return
	}

	//Если отсутствует правило повторения - удаляем задачу, в противном случае расчитываем новую дату
	if t.Repeat == "" {
		err := h.repo.DeleteTask(idInt)
		if err != nil {
			error.JsonResponse(w, "Ошибка удаления задачи")
			return
		}
	} else {
		now := time.Now()
		nextDate, err := nextdate.NextDate(now, t.Date, t.Repeat)
		if err != nil {
			error.JsonResponse(w, "Неверный формат правила повторения")
			return
		}
		//Update запрос к db
		res, err := h.repo.UpdateTask(nextDate, t.Title, t.Comment, t.Repeat, idInt)
		if err != nil || res == 0 {
			error.JsonResponse(w, "Ошибка обновления задачи")
			return
		}

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{}"))
}
