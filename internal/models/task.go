package models

import (
	"time"
)

type Task struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Task        string     `json:"task"`
	Description string     `gorm:"column:task" json:"description"`
	IsDone      bool       `json:"is_done"`
	UserID      uint       `json:"user_id"`                                 // ДОБАВЛЕНО: связь с пользователем
	User        User       `gorm:"foreignKey:UserID" json:"user,omitempty"` // ДОБАВЛЕНО: связь
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type TaskRequest struct {
	IsDone bool   `json:"is_done"`
	Task   string `json:"task"`
	UserID uint   `json:"user_id"` // ДОБАВЛЕНО: ID пользователя
}

func NewTask(req TaskRequest) *Task {
	return &Task{
		Description: req.Task,
		IsDone:      req.IsDone,
		UserID:      req.UserID, // ДОБАВЛЕНО
	}
}
