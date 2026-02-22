package models

// Status represents a Kanban column belonging to a Project.
// Maps to the `statuses` table.
type Status struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProjectID uint   `gorm:"not null;index" json:"projectId"`
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
	Color     string `gorm:"type:varchar(7);default:'#94A3B8'" json:"color"`
	Order     int    `gorm:"column:order;not null;default:0" json:"order"`
}

// TableName overrides the default GORM table name.
func (Status) TableName() string {
	return "statuses"
}
