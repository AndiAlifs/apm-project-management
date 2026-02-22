package models

// SubTask represents a checklist item belonging to a Task.
// Maps to the `sub_tasks` table.
type SubTask struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	TaskID uint   `gorm:"not null" json:"taskId"`
	Title  string `gorm:"type:varchar(500);not null" json:"title"`
	IsDone bool   `gorm:"default:false" json:"isDone"`
	Order  int    `gorm:"column:order;not null;default:0" json:"order"`
}

// TableName overrides the default GORM table name.
func (SubTask) TableName() string {
	return "sub_tasks"
}
