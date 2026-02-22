package services

import (
	"apm/backend/internal/models"

	"gorm.io/gorm"
)

// SubTaskService handles database operations for sub-tasks.
type SubTaskService struct {
	db *gorm.DB
}

// NewSubTaskService creates a new SubTaskService.
func NewSubTaskService(db *gorm.DB) *SubTaskService {
	return &SubTaskService{db: db}
}

// Create inserts a new sub-task under a task.
func (s *SubTaskService) Create(subtask *models.SubTask) error {
	var maxOrder int
	s.db.Model(&models.SubTask{}).
		Where("task_id = ?", subtask.TaskID).
		Select("COALESCE(MAX(`order`), -1)").
		Scan(&maxOrder)
	subtask.Order = maxOrder + 1
	return s.db.Create(subtask).Error
}

// Update saves changes to a sub-task (title, is_done, order).
func (s *SubTaskService) Update(id uint, updates *models.SubTask) (*models.SubTask, error) {
	var subtask models.SubTask
	if err := s.db.First(&subtask, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&subtask).Updates(updates).Error; err != nil {
		return nil, err
	}
	// Re-fetch to get updated values including booleans
	if err := s.db.First(&subtask, id).Error; err != nil {
		return nil, err
	}
	return &subtask, nil
}

// Delete removes a sub-task by ID.
func (s *SubTaskService) Delete(id uint) error {
	return s.db.Delete(&models.SubTask{}, id).Error
}
