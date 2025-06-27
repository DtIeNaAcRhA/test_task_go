package main

import (
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	tasks   = make(map[string]*Task)
	tasksMu sync.RWMutex
)

func main() {
	http.HandleFunc("/task", handleCreateTask)
	http.HandleFunc("/task/", handleTaskByID)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func processTask(id string) {
	tasksMu.Lock()
	task, exists := tasks[id]
	if !exists {
		tasksMu.Unlock()
		return
	}
	task.Status = Running
	tasksMu.Unlock()
	start := time.Now()

	duration := time.Duration(rand.Intn(3)+3) * time.Minute
	time.Sleep(duration)

	tasksMu.Lock()
	task.Status = Completed
	task.Duration = time.Since(start).String()
	task.Result = "done"
	tasksMu.Unlock()
}
