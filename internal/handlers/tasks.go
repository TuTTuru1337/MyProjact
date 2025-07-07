package handlers

import (
	"Tutturu/internal/models"
	"Tutturu/internal/repository"
	"encoding/json"
	"net/http"
	"strings"
)

type TaskHandler struct {
	repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	respondJSON(w, tasks)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	task, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	respondJSON(w, task)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Task == "" {
		http.Error(w, "Task description is required", http.StatusBadRequest)
		return
	}

	task := models.NewTask(req)
	if err := h.repo.Create(task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	respondJSON(w, task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	task, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var req models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Task != "" {
		task.Description = req.Task
	}
	task.IsDone = req.IsDone

	if err := h.repo.Update(task); err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	respondJSON(w, task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
