package service

import (
	"Tutturu/internal/models"
	"Tutturu/internal/userService/repository"
	"context"
	"strconv"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetTasksForUser(ctx context.Context, userID uint) ([]models.Task, error) // ДОБАВЛЕНО: новый метод
}

type UserServiceImpl struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserServiceImpl) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, strconv.Itoa(id))
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	// Здесь можно добавить хеширование пароля
	return s.repo.Create(ctx, user)
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return s.repo.Update(ctx, user)
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, strconv.Itoa(id))
}

// ДОБАВЛЕНО: новый метод для получения задач пользователя
func (s *UserServiceImpl) GetTasksForUser(ctx context.Context, userID uint) ([]models.Task, error) {
	return s.repo.GetTasksForUser(ctx, userID)
}
