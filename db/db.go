package db

import (
	"OneProject/internal/TaskService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	if err != nil {
	}
	if err := db.AutoMigrate(&TaskService.RequestBody{}); err != nil {
	}
	return db, nil
}
