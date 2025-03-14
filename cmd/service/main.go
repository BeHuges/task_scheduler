package main

import (
	"log"
	"net/http"
	"os"

	"task_scheduler/internal/database"
	"task_scheduler/internal/handler"
	"task_scheduler/internal/middleware"
	"task_scheduler/internal/repository"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("TODO_PORT")

	db := database.New()

	repo := repository.New(db)

	database.Migration(repo)

	handler := handler.New(repo)

	r := chi.NewRouter()

	r.Use(middleware.Logging)

	r.Handle("/*", http.FileServer(http.Dir("./web")))
	r.Get("/api/nextdate", handler.NextDate)
	r.Post("/api/task", handler.AddTask)
	r.Get("/api/tasks", handler.GetTasks)
	r.Get("/api/task", handler.GetTaskById)
	r.Put("/api/task", handler.UpdateTask)
	r.Post("/api/task/done", handler.TaskDone)
	r.Delete("/api/task", handler.DeleteTask)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
