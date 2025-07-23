package main

import (
	"Tutturu/internal/handlers"
	"Tutturu/internal/repository"
	"Tutturu/pkg/config"
	"Tutturu/pkg/database"
	"log"
	"net/http"
)

func main() {

	cfg := config.Load()

	db, err := database.InitDB(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	taskRepo := repository.NewTaskRepository(db)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetTasks(w, r)
		case http.MethodPost:
			taskHandler.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetTask(w, r)
		case http.MethodPatch:
			taskHandler.UpdateTask(w, r)
		case http.MethodDelete:
			taskHandler.DeleteTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Printf("Starting server on %s", cfg.Server.Address)
	if err := http.ListenAndServe(cfg.Server.Address, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
