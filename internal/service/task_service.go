package service

import (
	"Tutturu/internal/models"
	"Tutturu/internal/repository"
	"context"
)

type TaskService interface {
	GetAllTasks(ctx context.Context) ([]models.Task, error)
	CreateTask(ctx context.Context, task models.Task) (models.Task, error)
	GetTaskByID(ctx context.Context, id int) (*models.Task, error)
	UpdateTask(ctx context.Context, task models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, id int) error
	GetTasksByUserID(ctx context.Context, userID uint) ([]models.Task, error)
}

type Service struct {
	repo *repository.TaskRepository
}

func NewService(repo *repository.TaskRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	return s.repo.Create(ctx, task)
}

func (s *Service) GetTaskByID(ctx context.Context, id int) (*models.Task, error) {
	return s.repo.GetByID(ctx, uint(id))
}

func (s *Service) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
	return s.repo.Update(ctx, task)
}

func (s *Service) DeleteTask(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, uint(id))
}

func (s *Service) GetTasksByUserID(ctx context.Context, userID uint) ([]models.Task, error) {
	return s.repo.GetByUserID(ctx, userID)
}
