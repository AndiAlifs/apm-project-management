package services

import (
	"time"

	"apm/backend/internal/models"

	"gorm.io/gorm"
)

// TaskFilter holds optional query parameters for listing tasks.
type TaskFilter struct {
	ProjectID *uint
	StatusID  *uint
	Priority  *string
	TagID     *uint
	DueFrom   *time.Time
	DueTo     *time.Time
}

// TaskService handles all database operations for Tasks.
type TaskService struct {
	db *gorm.DB
}

// NewTaskService creates a new TaskService.
func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

// List returns tasks with optional filters, including associations.
func (s *TaskService) List(f TaskFilter) ([]models.Task, error) {
	query := s.db.
		Preload("Project").
		Preload("Status").
		Preload("SubTasks", func(db *gorm.DB) *gorm.DB {
			return db.Order("sub_tasks.order ASC")
		}).
		Preload("Tags")

	if f.ProjectID != nil {
		query = query.Where("tasks.project_id = ?", *f.ProjectID)
	}
	if f.StatusID != nil {
		query = query.Where("tasks.status_id = ?", *f.StatusID)
	}
	if f.Priority != nil {
		query = query.Where("tasks.priority = ?", *f.Priority)
	}
	if f.TagID != nil {
		query = query.Joins("JOIN task_tags ON task_tags.task_id = tasks.id").
			Where("task_tags.tag_id = ?", *f.TagID)
	}
	if f.DueFrom != nil {
		query = query.Where("tasks.due_date >= ?", *f.DueFrom)
	}
	if f.DueTo != nil {
		query = query.Where("tasks.due_date <= ?", *f.DueTo)
	}

	var tasks []models.Task
	err := query.Order("tasks.created_at DESC").Find(&tasks).Error
	return tasks, err
}

// ListByProject returns tasks for a single project, grouped by status.
func (s *TaskService) ListByProject(projectID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := s.db.
		Preload("Status").
		Preload("SubTasks", func(db *gorm.DB) *gorm.DB {
			return db.Order("sub_tasks.order ASC")
		}).
		Preload("Tags").
		Where("project_id = ?", projectID).
		Find(&tasks).Error
	return tasks, err
}

// GetByID returns a fully preloaded task.
func (s *TaskService) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := s.db.
		Preload("Project").
		Preload("Status").
		Preload("SubTasks", func(db *gorm.DB) *gorm.DB {
			return db.Order("sub_tasks.order ASC")
		}).
		Preload("Tags").
		First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Create inserts a new task.
func (s *TaskService) Create(task *models.Task) error {
	return s.db.Create(task).Error
}

// Update saves changes to a task. Handles started_at / completed_at transitions.
func (s *TaskService) Update(id uint, updates map[string]interface{}) (*models.Task, error) {
	var task models.Task
	if err := s.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&task).Updates(updates).Error; err != nil {
		return nil, err
	}
	return s.GetByID(id)
}

// PatchStatus moves a task to a new status column.
// Auto-sets started_at and completed_at based on status name.
func (s *TaskService) PatchStatus(taskID, newStatusID uint) (*models.Task, error) {
	var task models.Task
	if err := s.db.Preload("Status").First(&task, taskID).Error; err != nil {
		return nil, err
	}

	var newStatus models.Status
	if err := s.db.First(&newStatus, newStatusID).Error; err != nil {
		return nil, err
	}

	updates := map[string]interface{}{"status_id": newStatusID}

	now := time.Now()
	if task.StartedAt == nil && newStatus.Name != "To Do" {
		updates["started_at"] = now
	}
	if newStatus.Name == "Done" && task.CompletedAt == nil {
		updates["completed_at"] = now
	} else if newStatus.Name != "Done" {
		updates["completed_at"] = nil
	}

	if err := s.db.Model(&task).Updates(updates).Error; err != nil {
		return nil, err
	}
	return s.GetByID(taskID)
}

// Delete removes a task by ID.
func (s *TaskService) Delete(id uint) error {
	return s.db.Delete(&models.Task{}, id).Error
}
