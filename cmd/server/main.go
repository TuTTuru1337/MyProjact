package main

import (
	"Tutturu/internal/handlers"
	"Tutturu/internal/models"
	"Tutturu/internal/repository"
	"Tutturu/internal/service"
	"Tutturu/internal/web/tasks"
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

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	repo := repository.NewTaskRepository(db)
	svc := service.NewService(repo)
	handler := handlers.NewHandler(svc)

	// Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(cfg.Server.Address); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
