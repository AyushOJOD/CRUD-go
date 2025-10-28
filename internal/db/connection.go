package db

import (
	"fmt"
	"log"

	"github.com/AyushOJOD/task-manager-api/config"
	"github.com/AyushOJOD/task-manager-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Could not connect to the database: ", err)
	}

	DB = conn
	log.Println("Connected to the database successfully")


	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Panic("Could not migrate the database: ", err)
	}

	log.Println("Database migrated successfully")
}