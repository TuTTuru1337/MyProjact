package service

import (
	"Tutturu/internal/models"
	"Tutturu/internal/repository"
	"context"
	"strconv"
)

type TaskService interface {
	GetAllTasks(ctx context.Context) ([]models.Task, error)
	CreateTask(ctx context.Context, task models.Task) (models.Task, error)
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
	createdTask, err := s.repo.Create(ctx, task)
	if err != nil {
		return models.Task{}, err
	}
	return createdTask, nil
}

func (s *Service) GetTaskByID(ctx context.Context, id int) (*models.Task, error) {
	return s.repo.GetByID(ctx, strconv.Itoa(id))
}

func (s *Service) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
	return s.repo.Update(ctx, task)
}
func (s *Service) DeleteTask(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, strconv.Itoa(id))
}
