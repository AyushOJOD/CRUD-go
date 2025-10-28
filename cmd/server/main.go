package main

import (
	"github.com/AyushOJOD/task-manager-api/config"
	"github.com/AyushOJOD/task-manager-api/internal/db"
	"github.com/AyushOJOD/task-manager-api/internal/handlers"
	"github.com/AyushOJOD/task-manager-api/internal/routes"
	"github.com/AyushOJOD/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config.LoadConfig()

	// Connect to DB
	db.ConnectDB()

	// Initialize Gin router
	r := gin.Default()

	// Initialize service and handler
	taskService := services.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)

    // Register routes via routes package
    routes.SetupRoutes(r, taskHandler)

	// Start server
	r.Run(":8080")
}