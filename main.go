package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

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

func main() {
	http.HandleFunc("/task", postHandler)
	http.HandleFunc("/get", GetHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
