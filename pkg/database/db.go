package database

import (
	"Tutturu/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return nil, fmt.Errorf("ошибка при миграции базы данных: %w", err)
	}

	return db, nil
}
