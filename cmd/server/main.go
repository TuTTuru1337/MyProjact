package main

import (
	"Tutturu/internal/handlers"
	"Tutturu/internal/models"
	"Tutturu/internal/repository"
	"Tutturu/internal/service"
	userRepo "Tutturu/internal/userService/repository"
	userService "Tutturu/internal/userService/service"
	"Tutturu/internal/web/tasks"
	"Tutturu/internal/web/users"
	"Tutturu/pkg/config"
	"Tutturu/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"strings"
)

func main() {
	cfg := config.Load()
	db, err := database.InitDB(cfg.DB.DSN)
	if err != nil {
		if strings.Contains(err.Error(), "миграции") {
			log.Fatalf("Ошибка миграции базы данных: %v", err)
		} else {
			log.Fatalf("Ошибка подключения к базе данных: %v", err)
		}
	}

	if err := db.AutoMigrate(&models.Task{}, &models.User{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	tasksRepo := repository.NewTaskRepository(db)
	tasksService := service.NewService(tasksRepo)
	tasksHandler := handlers.NewHandler(tasksService)

	userRepository := userRepo.NewUserRepository(db)
	userServiceImpl := userService.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userServiceImpl)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	usersStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(cfg.Server.Address); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
