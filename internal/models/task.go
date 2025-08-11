package models

import (
	_ "gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Task        string     `json:"task"`
	Description string     `gorm:"column:task" json:"description"` // Указываем имя колонки task
	IsDone      bool       `json:"is_done"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type TaskRequest struct {
	IsDone bool   `json:"is_done"`
	Task   string `json:"task"`
}

func NewTask(req TaskRequest) *Task {
	return &Task{
		Description: req.Task,
		IsDone:      req.IsDone,
	}
}
