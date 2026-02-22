package services

import (
	"apm/backend/internal/models"

	"gorm.io/gorm"
)

// TagService handles database operations for Tags and task-tag associations.
type TagService struct {
	db *gorm.DB
}

// NewTagService creates a new TagService.
func NewTagService(db *gorm.DB) *TagService {
	return &TagService{db: db}
}

// List returns all tags ordered by name.
func (s *TagService) List() ([]models.Tag, error) {
	var tags []models.Tag
	err := s.db.Order("name ASC").Find(&tags).Error
	return tags, err
}

// Create inserts a new tag.
func (s *TagService) Create(tag *models.Tag) error {
	return s.db.Create(tag).Error
}

// Update saves changes to a tag (name, colour).
func (s *TagService) Update(id uint, updates *models.Tag) (*models.Tag, error) {
	var tag models.Tag
	if err := s.db.First(&tag, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&tag).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

// Delete removes a tag by ID (cascades task_tags via DB constraint).
func (s *TagService) Delete(id uint) error {
	return s.db.Delete(&models.Tag{}, id).Error
}

// AssignToTask creates a task-tag association.
func (s *TagService) AssignToTask(taskID, tagID uint) error {
	var task models.Task
	if err := s.db.First(&task, taskID).Error; err != nil {
		return err
	}
	var tag models.Tag
	if err := s.db.First(&tag, tagID).Error; err != nil {
		return err
	}
	return s.db.Model(&task).Association("Tags").Append(&tag)
}

// RemoveFromTask deletes a task-tag association.
func (s *TagService) RemoveFromTask(taskID, tagID uint) error {
	var task models.Task
	if err := s.db.First(&task, taskID).Error; err != nil {
		return err
	}
	var tag models.Tag
	if err := s.db.First(&tag, tagID).Error; err != nil {
		return err
	}
	return s.db.Model(&task).Association("Tags").Delete(&tag)
}
