package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}

func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	log.Println("Database migration completed")
	return nil
}
