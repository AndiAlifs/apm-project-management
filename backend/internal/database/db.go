package database

import (
	"fmt"
	"log"

	"apm/backend/internal/config"
	"apm/backend/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the package-level GORM database instance.
var DB *gorm.DB

// Connect opens a MySQL connection using the provided config and stores it in DB.
// It also runs AutoMigrate to ensure all tables are up to date with the model definitions.
func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")

	// AutoMigrate ensures all tables exist and schema is up to date.
	// Tables: projects, statuses, tasks, sub_tasks, tags, task_tags
	err = db.AutoMigrate(
		&models.Project{},
		&models.Status{},
		&models.Task{},
		&models.SubTask{},
		&models.Tag{},
		&models.TaskTag{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	log.Println("Database schema migrated successfully")
	DB = db
}
