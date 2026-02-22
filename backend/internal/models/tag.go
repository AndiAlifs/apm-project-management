package models

// Tag is a global, cross-cutting label that can be assigned to many Tasks.
// Maps to the `tags` table.
type Tag struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:varchar(100);not null;uniqueIndex" json:"name"`
	Color string `gorm:"type:varchar(7);default:'#F59E0B'" json:"color"`

	// Many-to-many back-reference (not stored in this table)
	Tasks []Task `gorm:"many2many:task_tags;" json:"tasks,omitempty"`
}

// TaskTag is the explicit join table for the Task <-> Tag many-to-many relationship.
// Maps to the `task_tags` table.
// GORM uses this automatically via the `many2many:"task_tags"` tag; we define it
// explicitly so AutoMigrate creates the composite PK correctly.
type TaskTag struct {
	TaskID uint `gorm:"primaryKey" json:"taskId"`
	TagID  uint `gorm:"primaryKey" json:"tagId"`
}

// TableName overrides the default GORM table name.
func (TaskTag) TableName() string {
	return "task_tags"
}
