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

	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = database.Migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	taskRepo := repository.NewTaskRepository(db)
	userRepo := userrepository.NewUserRepository(db)
	taskService := service.NewService(taskRepo)
	userServiceImpl := userservice.NewUserService(userRepo) // УБРАТЬ ЛИШНИЙ ПАРАМЕТР
	taskHandler := handlers.NewHandler(taskService)
	userHandler := handlers.NewUserHandler(userServiceImpl)

	e := echo.New()
	tasks.RegisterHandlers(e, taskHandler)
	users.RegisterHandlers(e, userHandler)

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(e.Start(":" + cfg.Port))
}
