package routes

import (
	"net/http"

	"github.com/AyushOJOD/task-manager-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers all HTTP routes on the given Gin engine.
func SetupRoutes(r *gin.Engine, taskHandler *handlers.TaskHandler) {
    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    // API Group
    api := r.Group("/api")
    {
        tasks := api.Group("/tasks")
        {
            tasks.POST("", taskHandler.CreateTask)
            tasks.GET("", taskHandler.GetTasks)
            tasks.GET("/:id", taskHandler.GetTask)
            tasks.PUT("/:id", taskHandler.UpdateTask)
            tasks.DELETE("/:id", taskHandler.DeleteTask)
        }
    }
}


