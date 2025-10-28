package services

import (
	"github.com/AyushOJOD/task-manager-api/internal/db"
	"github.com/AyushOJOD/task-manager-api/internal/models"
	"gorm.io/gorm"
)

type TaskService struct {
    DB *gorm.DB  // We'll inject the global DB
}

// NewTaskService creates a new service instance
func NewTaskService() *TaskService {
    return &TaskService{DB: db.DB}
}

// Create a new task
func (s *TaskService) Create(task *models.Task) error {
    return s.DB.Create(task).Error
}

// Get all tasks
func (s *TaskService) GetAll() ([]models.Task, error) {
    var tasks []models.Task
    err := s.DB.Find(&tasks).Error
    return tasks, err
}

// Get a task by ID
func (s *TaskService) GetByID(id uint) (*models.Task, error) {
    var task models.Task
    err := s.DB.First(&task, id).Error
    if err != nil {
        return nil, err
    }
    return &task, nil
}

// Update a task by ID
func (s *TaskService) Update(id uint, task *models.Task) error {
    // Fetch the existing task to avoid overwriting with zeros
    existingTask, err := s.GetByID(id)
    if err != nil {
        return err
    }

    // Update fields
    existingTask.Title = task.Title
    existingTask.Description = task.Description
    existingTask.Completed = task.Completed

    return s.DB.Save(existingTask).Error
}

// Delete a task by ID (soft delete via GORM)
func (s *TaskService) Delete(id uint) error {
    var task models.Task
    err := s.DB.First(&task, id).Error
    if err != nil {
        return err
    }
    return s.DB.Delete(&task).Error
}