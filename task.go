package main

import (
	"time"
)

type TaskStatus string

const (
	Pending   TaskStatus = "pending"   //в ожидании
	Running   TaskStatus = "running"   //выполняется
	Completed TaskStatus = "completed" //завершена
)

type Task struct {
	ID        string     `json:"id"`               //id (uuid)
	CreatedAt time.Time  `json:"created_at"`       //дата создания (текущая на момент создания)
	Status    TaskStatus `json:"status"`           //статус
	Duration  string     `json:"duration"`         //продолжительность
	Result    string     `json:"result,omitempty"` //результат задчи
}
