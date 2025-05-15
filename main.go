package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"sync"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

var (
	tasks    = make(map[string]Task)
	tasksMux sync.RWMutex
)

type TaskRequest struct {
	Description string `json:"description"`
}

func generateID() string {
	return uuid.New().String()
}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid requst method", http.StatusMethodNotAllowed)
		return
	}
	tasksMux.RLock()
	defer tasksMux.RUnlock()

	taskList := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskList)
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

	if req.Description == "" {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}

	newTask := Task{ID: generateID(), Description: req.Description}

	tasksMux.Lock()
	tasks[newTask.ID] = newTask
	tasksMux.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Patch only", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Path[len("/patch/"):]

	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tasksMux.Lock()
	defer tasksMux.Unlock()

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if task.Description != "" {
		task.Description = req.Description
		tasks[id] = task
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	tasksMux.Lock()
	defer tasksMux.Unlock()

	if _, ok := tasks[id]; !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	delete(tasks, id)
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

func main() {
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
