package main

import (
	"Tutturu/internal/handlers"
	"Tutturu/internal/pkg/config"
	"Tutturu/internal/pkg/database"
	"Tutturu/internal/repository"
	"Tutturu/internal/service"
	userrepository "Tutturu/internal/userService/repository"
	userservice "Tutturu/internal/userService/service"
	"Tutturu/internal/web/tasks"
	"Tutturu/internal/web/users"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Загрузка конфигурации
	cfg := config.Load()

	// Подключение к базе данных
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Автомиграция (опционально)
	err = database.Migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Инициализация репозиториев
	taskRepo := repository.NewTaskRepository(db)
	userRepo := userrepository.NewUserRepository(db)

	// Инициализация сервисов
	taskService := service.NewService(taskRepo)
	userServiceImpl := userservice.NewUserService(userRepo) // УБРАТЬ ЛИШНИЙ ПАРАМЕТР

	// Инициализация хендлеров
	taskHandler := handlers.NewHandler(taskService)
	userHandler := handlers.NewUserHandler(userServiceImpl)

	// Создание Echo router
	e := echo.New()

	// Регистрация роутов для задач
	tasks.RegisterHandlers(e, taskHandler)

	// Регистрация роутов для пользователей
	users.RegisterHandlers(e, userHandler)

	// Запуск сервера
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(e.Start(":" + cfg.Port))
}
