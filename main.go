package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var Task string

type TaskRequest struct {
	Task string `json:"task"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post only", http.StatusMethodNotAllowed)
		return
	}

	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	Task = req.Task
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task updated to: %s", Task)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Current task: %s", Task)
}

var (
	tasks    = make(map[string]string)
	tasksMux sync.RWMutex
)

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Patch only", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID   string `json:"id"`
		Task string `json:"task"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tasksMux.Lock()
	defer tasksMux.Unlock()

	tasks[req.ID] = req.Task
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task '%s' updated", req.ID)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Delete only", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	tasksMux.Lock()
	defer tasksMux.Unlock()

	delete(tasks, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task '%s' deleted", id)
}

func main() {

	http.HandleFunc("/task", postHandler)
	http.HandleFunc("/get", GetHandler)

	http.HandleFunc("/task/patch", PatchHandler)
	http.HandleFunc("/task/delete", DeleteHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
