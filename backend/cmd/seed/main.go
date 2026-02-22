package main

import (
	"log"
	"time"

	"apm/backend/internal/config"
	"apm/backend/internal/database"
	"apm/backend/internal/models"
)

func main() {
	cfg := config.Load()
	database.Connect(cfg)
	db := database.DB

	log.Println("Starting seed...")

	// ── Tags ──────────────────────────────────────────────────────────────────
	tags := []models.Tag{
		{Name: "job", Color: "#3B82F6"},
		{Name: "freelance", Color: "#8B5CF6"},
		{Name: "personal", Color: "#22C55E"},
	}
	for i := range tags {
		if err := db.FirstOrCreate(&tags[i], models.Tag{Name: tags[i].Name}).Error; err != nil {
			log.Fatalf("Failed to seed tag %q: %v", tags[i].Name, err)
		}
	}
	log.Printf("  ✓ Tags seeded: %d records", len(tags))

	// ── Project ───────────────────────────────────────────────────────────────
	project := models.Project{
		Name:        "My First Project",
		Description: "Sample project seeded for Sprint 02 testing",
		Color:       "#6366F1",
	}
	if err := db.FirstOrCreate(&project, models.Project{Name: project.Name}).Error; err != nil {
		log.Fatalf("Failed to seed project: %v", err)
	}
	log.Printf("  ✓ Project seeded: id=%d", project.ID)

	// Seed default statuses if they don't exist
	var statusCount int64
	db.Model(&models.Status{}).Where("project_id = ?", project.ID).Count(&statusCount)
	if statusCount == 0 {
		defaultStatuses := []models.Status{
			{ProjectID: project.ID, Name: "To Do", Color: "#94A3B8", Order: 0},
			{ProjectID: project.ID, Name: "In Progress", Color: "#3B82F6", Order: 1},
			{ProjectID: project.ID, Name: "Done", Color: "#22C55E", Order: 2},
		}
		if err := db.Create(&defaultStatuses).Error; err != nil {
			log.Fatalf("Failed to seed statuses: %v", err)
		}
	}

	// Fetch statuses
	var statuses []models.Status
	if err := db.Where("project_id = ?", project.ID).Order("`order` ASC").Find(&statuses).Error; err != nil {
		log.Fatalf("Failed to fetch statuses: %v", err)
	}
	log.Printf("  ✓ Statuses fetched: %d records", len(statuses))

	todoStatus := statuses[0]
	inProgressStatus := statuses[1]
	doneStatus := statuses[2]

	// ── Tasks ─────────────────────────────────────────────────────────────────
	due1 := time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC)
	due2 := time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC)
	due3 := time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC)
	due4 := time.Date(2026, 5, 31, 0, 0, 0, 0, time.UTC)

	now := time.Now()
	tasks := []models.Task{
		{
			ProjectID:   project.ID,
			StatusID:    todoStatus.ID,
			Title:       "Set up CI/CD pipeline",
			Description: "Configure GitHub Actions for build, test, and deploy steps.",
			Priority:    models.PriorityHigh,
			DueDate:     &due1,
		},
		{
			ProjectID:   project.ID,
			StatusID:    todoStatus.ID,
			Title:       "Write API documentation",
			Description: "Document all REST endpoints using Swagger/OpenAPI.",
			Priority:    models.PriorityMedium,
			DueDate:     &due2,
		},
		{
			ProjectID:   project.ID,
			StatusID:    inProgressStatus.ID,
			Title:       "Build Angular app shell",
			Description: "Scaffold the Angular 18 frontend with Tailwind CSS, sidebar, and top navigation.",
			Priority:    models.PriorityHigh,
			DueDate:     &due3,
			StartedAt:   &now,
		},
		{
			ProjectID:   project.ID,
			StatusID:    inProgressStatus.ID,
			Title:       "Design database schema",
			Description: "Create ERD and GORM models for all core entities.",
			Priority:    models.PriorityUrgent,
			DueDate:     &due3,
			StartedAt:   &now,
		},
		{
			ProjectID:   project.ID,
			StatusID:    doneStatus.ID,
			Title:       "Initialize Go backend",
			Description: "Set up Go module, Gin router, GORM, and .env config.",
			Priority:    models.PriorityMedium,
			DueDate:     &due4,
			StartedAt:   &now,
			CompletedAt: &now,
		},
	}

	for i := range tasks {
		if err := db.FirstOrCreate(&tasks[i], models.Task{Title: tasks[i].Title, ProjectID: tasks[i].ProjectID}).Error; err != nil {
			log.Fatalf("Failed to seed task %q: %v", tasks[i].Title, err)
		}
	}
	log.Printf("  ✓ Tasks seeded: %d records", len(tasks))

	// ── Sub-tasks ─────────────────────────────────────────────────────────────
	subTasks := []models.SubTask{
		{TaskID: tasks[0].ID, Title: "Create .github/workflows directory", IsDone: true, Order: 0},
		{TaskID: tasks[0].ID, Title: "Add build and test job", IsDone: false, Order: 1},
		{TaskID: tasks[0].ID, Title: "Add deployment job", IsDone: false, Order: 2},
		{TaskID: tasks[2].ID, Title: "Install Angular CLI", IsDone: true, Order: 0},
		{TaskID: tasks[2].ID, Title: "Configure Tailwind CSS", IsDone: false, Order: 1},
	}
	for i := range subTasks {
		if err := db.FirstOrCreate(&subTasks[i], models.SubTask{TaskID: subTasks[i].TaskID, Title: subTasks[i].Title}).Error; err != nil {
			log.Fatalf("Failed to seed subtask: %v", err)
		}
	}
	log.Printf("  ✓ Sub-tasks seeded: %d records", len(subTasks))

	// ── Tag Assignments ───────────────────────────────────────────────────────
	jobTag := tags[0]
	freelanceTag := tags[1]
	personalTag := tags[2]

	assignments := []struct {
		task *models.Task
		tag  *models.Tag
	}{
		{&tasks[0], &jobTag},
		{&tasks[1], &jobTag},
		{&tasks[1], &freelanceTag},
		{&tasks[2], &freelanceTag},
		{&tasks[3], &jobTag},
		{&tasks[4], &personalTag},
	}
	for _, a := range assignments {
		if err := db.Model(a.task).Association("Tags").Append(a.tag); err != nil {
			log.Printf("  (warn) tag assignment skipped (may already exist): %v", err)
		}
	}
	log.Printf("  ✓ Tag assignments seeded: %d records", len(assignments))

	log.Println("Seed completed successfully! 🎉")
}
