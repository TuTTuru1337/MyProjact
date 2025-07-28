package handlers

import (
	"Tutturu/internal/models"
	"Tutturu/internal/repository"
	"encoding/json"
	"log"
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
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	tasks, err := h.repo.GetAll()
	if err != nil {
		log.Printf("Ошибка при получении задач: %v", err)
		respondError(w, http.StatusInternalServerError, "Не удалось получить список задач")
		return
	}

	respondJSON(w, tasks)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	task, err := h.repo.GetByID(id)
	if err != nil {
		log.Printf("Ошибка при получении задачи по id=%s: %v", id, err)
		respondError(w, http.StatusNotFound, "Задача не найдена")
		return
	}

	respondJSON(w, task)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	var req models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Невалидное тело запроса")
		return
	}
	defer r.Body.Close()

	if req.Task == "" {
		respondError(w, http.StatusBadRequest, "Описание задачи обязательно")
		return
	}

	task := models.NewTask(req)
	if err := h.repo.Create(task); err != nil {
		log.Printf("Ошибка при создании задачи: %v", err)
		respondError(w, http.StatusInternalServerError, "Не удалось создать задачу")
		return
	}

	w.WriteHeader(http.StatusCreated)
	respondJSON(w, task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	task, err := h.repo.GetByID(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Задача не найдена")
		return
	}

	var req models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Невалидное тело запроса")
		return
	}
	defer r.Body.Close()

	if req.Task != "" {
		task.Description = req.Task
	}
	task.IsDone = req.IsDone

	if err := h.repo.Update(task); err != nil {
		log.Printf("Ошибка при обновлении задачи id=%s: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Не удалось обновить задачу")
		return
	}

	respondJSON(w, task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "Метод не разрешён")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if err := h.repo.Delete(id); err != nil {
		log.Printf("Ошибка при удалении задачи id=%s: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Не удалось удалить задачу")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	log.Printf("HTTP %d - %s", status, message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
