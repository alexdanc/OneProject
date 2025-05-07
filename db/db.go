package db

import (
	"OneProject/internal/TaskService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	if err := db.AutoMigrate(&TaskService.RequestBody{}); err != nil {
		log.Fatal("Failed to migrate DB: ", err)
	}
	return db, nil
}
