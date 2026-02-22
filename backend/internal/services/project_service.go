package services

import (
	"apm/backend/internal/models"

	"gorm.io/gorm"
)

// ProjectService handles all database operations for Projects.
type ProjectService struct {
	db *gorm.DB
}

// NewProjectService creates a new ProjectService.
func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{db: db}
}

// List returns all projects ordered by creation date.
func (s *ProjectService) List() ([]models.Project, error) {
	var projects []models.Project
	err := s.db.Order("created_at ASC").Find(&projects).Error
	return projects, err
}

// GetByID returns a single project including its statuses.
func (s *ProjectService) GetByID(id uint) (*models.Project, error) {
	var project models.Project
	err := s.db.Preload("Statuses", func(db *gorm.DB) *gorm.DB {
		return db.Order("statuses.order ASC")
	}).First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// Create inserts a new project and seeds default Kanban statuses.
func (s *ProjectService) Create(p *models.Project) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(p).Error; err != nil {
			return err
		}
		// Seed default statuses for new projects
		defaultStatuses := []models.Status{
			{ProjectID: p.ID, Name: "To Do", Color: "#94A3B8", Order: 0},
			{ProjectID: p.ID, Name: "In Progress", Color: "#3B82F6", Order: 1},
			{ProjectID: p.ID, Name: "Done", Color: "#22C55E", Order: 2},
		}
		return tx.Create(&defaultStatuses).Error
	})
}

// Update saves changes to an existing project.
func (s *ProjectService) Update(id uint, updates *models.Project) (*models.Project, error) {
	var project models.Project
	if err := s.db.First(&project, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&project).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

// Delete removes a project (cascades to tasks via DB constraint).
func (s *ProjectService) Delete(id uint) error {
	return s.db.Delete(&models.Project{}, id).Error
}
