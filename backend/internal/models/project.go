package models

import "time"

// Project represents a top-level container for Tasks.
// Maps to the `projects` table.
type Project struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Color       string    `gorm:"type:varchar(7);default:'#6366F1'" json:"color"`
	Icon        string    `gorm:"type:varchar(50)" json:"icon"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	// Associations (used for eager loading, not stored in this table)
	Statuses []Status `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"statuses,omitempty"`
	Tasks    []Task   `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"tasks,omitempty"`
}
