package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=YOUPASSWORD dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

}

type Task struct {
	ID          string     `gorm:"primaryKey" json:"id"`
	Description string     `json:"task"`
	IsDone      bool       `json:"is_done"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

type TaskRequest struct {
	IsDone bool   `json:"is_Done"`
	Task   string `json:"task"`
}

func generateID() string {
	return uuid.New().String()
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid requst method", http.StatusMethodNotAllowed)
		return
	}

	var tasks []Task

	if err := db.Where("deleted_at IS NULL").Find(&tasks).Error; err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid requst method", http.StatusMethodNotAllowed)
		return
	}
	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Task == "" {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}

	newTask := Task{
		ID:          generateID(),
		Description: req.Task,
		IsDone:      req.IsDone,
	}
	if err := db.Create(&newTask).Error; err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
func PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")

	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var task Task
	if err := db.First(&task, "id = ? AND deleted_at IS NULL", id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.Description = req.Task
	task.IsDone = req.IsDone

	if err := db.Save(&task).Error; err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")

	var task Task
	if err := db.First(&task, "id = ? AND deleted_at IS NULL", id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	deletedAt := time.Now()
	task.DeletedAt = &deletedAt

	if err := db.Save(&task).Error; err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
func main() {

	initDB()

	http.HandleFunc("/tasks", GetTaskHandler)
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			PostTaskHandler(w, r)
		case http.MethodPatch:
			PatchTaskHandler(w, r)
		case http.MethodDelete:
			DeleteTaskHandler(w, r)
		case http.MethodGet:

			http.Error(w, "Not Implemented", http.StatusNotImplemented)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
