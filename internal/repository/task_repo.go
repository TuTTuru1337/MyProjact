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
	err := r.db.WithContext(ctx).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) Create(ctx context.Context, task models.Task) (models.Task, error) {
	err := r.db.WithContext(ctx).Create(&task).Error
	return task, err
}

func (r *TaskRepository) GetByID(ctx context.Context, id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(ctx context.Context, task models.Task) (models.Task, error) {
	err := r.db.WithContext(ctx).Save(&task).Error
	return task, err
}

func (r *TaskRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Task{}, id).Error
}

func (r *TaskRepository) GetByUserID(ctx context.Context, userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
