package models

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
	"time"
)

type Task struct {
	ID          string     `gorm:"primaryKey" json:"id"`
	Description string     `json:"description"`
	IsDone      bool       `json:"is_done"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

type TaskRequest struct {
	IsDone bool   `json:"is_Done"`
	Task   string `json:"task"`
}

func NewTask(req TaskRequest) *Task {
	return &Task{
		ID:          uuid.New().String(),
		Description: req.Task,
		IsDone:      req.IsDone,
	}
}
