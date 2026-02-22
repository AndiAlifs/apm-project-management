package models

import "time"

// Priority represents the urgency level of a Task.
type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
	PriorityUrgent Priority = "urgent"
)

// Task represents a unit of work within a Project.
// Maps to the `tasks` table.
type Task struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ProjectID   uint       `gorm:"not null;index" json:"projectId"`
	StatusID    uint       `gorm:"not null;index" json:"statusId"`
	Title       string     `gorm:"type:varchar(500);not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Priority    Priority   `gorm:"type:enum('low','medium','high','urgent');default:'medium'" json:"priority"`
	DueDate     *time.Time `gorm:"type:date;index" json:"dueDate"`
	StartedAt   *time.Time `json:"startedAt"`
	CompletedAt *time.Time `json:"completedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`

	// Associations
	Project  Project   `gorm:"foreignKey:ProjectID" json:"project,omitempty"`
	Status   Status    `gorm:"foreignKey:StatusID" json:"status,omitempty"`
	SubTasks []SubTask `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE" json:"subTasks,omitempty"`
	Tags     []Tag     `gorm:"many2many:task_tags;" json:"tags,omitempty"`
}
