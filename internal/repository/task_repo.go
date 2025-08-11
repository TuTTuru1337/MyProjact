package repository

import (
	"Tutturu/internal/models"
	"context"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id string) (*models.Task, error) {
	var task models.Task
	err := r.db.WithContext(ctx).First(&task, "id = ? AND deleted_at IS NULL", id).Error
	return &task, err
}

func (r *TaskRepository) Create(ctx context.Context, task models.Task) (models.Task, error) {
	if err := r.db.WithContext(ctx).Create(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) Update(ctx context.Context, task models.Task) (models.Task, error) {
	err := r.db.WithContext(ctx).Save(&task).Error
	return task, err
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {

	return r.db.WithContext(ctx).Delete(&models.Task{}, "id = ?", id).Error
}
