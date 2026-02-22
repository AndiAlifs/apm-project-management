package services

import (
	"apm/backend/internal/models"

	"gorm.io/gorm"
)

// StatusService handles database operations for Kanban column statuses.
type StatusService struct {
	db *gorm.DB
}

// NewStatusService creates a new StatusService.
func NewStatusService(db *gorm.DB) *StatusService {
	return &StatusService{db: db}
}

// ListByProject returns all statuses for a project ordered by their display order.
func (s *StatusService) ListByProject(projectID uint) ([]models.Status, error) {
	var statuses []models.Status
	err := s.db.Where("project_id = ?", projectID).Order("`order` ASC").Find(&statuses).Error
	return statuses, err
}

// Create inserts a new status for a project.
func (s *StatusService) Create(status *models.Status) error {
	// Auto-set order to be last in the project
	var maxOrder int
	s.db.Model(&models.Status{}).
		Where("project_id = ?", status.ProjectID).
		Select("COALESCE(MAX(`order`), -1)").
		Scan(&maxOrder)
	status.Order = maxOrder + 1
	return s.db.Create(status).Error
}

// Update saves changes to a status (name, colour, order).
func (s *StatusService) Update(id uint, updates *models.Status) (*models.Status, error) {
	var status models.Status
	if err := s.db.First(&status, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&status).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &status, nil
}

// Delete removes a status by ID.
func (s *StatusService) Delete(id uint) error {
	return s.db.Delete(&models.Status{}, id).Error
}
