package repository

import (
	"Tutturu/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("deleted_at IS NULL").Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetByID(id string) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, "id = ? AND deleted_at IS NULL", id).Error
	return &task, err
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) Delete(id string) error {
	return r.db.Delete(&models.Task{}, "id = ?", id).Error
}
