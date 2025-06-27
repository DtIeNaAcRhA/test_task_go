package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	taskID := uuid.New().String()
	task := &Task{
		ID:        taskID,
		CreatedAt: time.Now(),
		Status:    Pending,
	}

	tasksMu.Lock()
	tasks[taskID] = task
	tasksMu.Unlock()

	go processTask(taskID)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("task  \"%s\" is created", taskID)))
	json.NewEncoder(w).Encode(task)
}

func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/task/"):]
	tasksMu.RLock()
	task, exists := tasks[id]
	tasksMu.RUnlock()

	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		start := time.Now()
		tasksMu.Lock()
		task.Duration = start.Sub(task.CreatedAt).String()
		tasksMu.Unlock()
		json.NewEncoder(w).Encode(task)
	case http.MethodDelete:
		tasksMu.Lock()
		delete(tasks, id)
		tasksMu.Unlock()
		w.Write([]byte(fmt.Sprintf("task  \"%s\" is deleted", id)))
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
