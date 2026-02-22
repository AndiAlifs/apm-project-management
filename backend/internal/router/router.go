package router

import (
	"net/http"

	"apm/backend/internal/config"
	"apm/backend/internal/database"
	"apm/backend/internal/handlers"
	"apm/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// Setup creates and configures the Gin engine with CORS middleware and all routes.
func Setup(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// ── CORS Middleware ────────────────────────────────────────────────────────
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", cfg.CORSOrigin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// ── Health Check ──────────────────────────────────────────────────────────
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// ── Services ─────────────────────────────────────────────────────────────
	db := database.DB
	projectSvc := services.NewProjectService(db)
	statusSvc := services.NewStatusService(db)
	taskSvc := services.NewTaskService(db)
	subtaskSvc := services.NewSubTaskService(db)
	tagSvc := services.NewTagService(db)

	// ── Handlers ─────────────────────────────────────────────────────────────
	projectH := handlers.NewProjectHandler(projectSvc)
	statusH := handlers.NewStatusHandler(statusSvc)
	taskH := handlers.NewTaskHandler(taskSvc)
	subtaskH := handlers.NewSubTaskHandler(subtaskSvc)
	tagH := handlers.NewTagHandler(tagSvc)

	// ── API v1 Routes ─────────────────────────────────────────────────────────
	api := r.Group("/api")
	{
		// Projects
		api.GET("/projects", projectH.List)
		api.POST("/projects", projectH.Create)
		api.PUT("/projects/:id", projectH.Update)
		api.DELETE("/projects/:id", projectH.Delete)

		// Statuses (Kanban columns) — nested under project
		api.GET("/projects/:id/statuses", statusH.ListByProject)
		api.POST("/projects/:id/statuses", statusH.Create)
		api.PUT("/statuses/:id", statusH.Update)
		api.DELETE("/statuses/:id", statusH.Delete)

		// Tasks — per-project and global backlog
		api.GET("/projects/:id/tasks", taskH.ListByProject)
		api.GET("/tasks", taskH.List)
		api.GET("/tasks/:id", taskH.GetByID)
		api.POST("/tasks", taskH.Create)
		api.PUT("/tasks/:id", taskH.Update)
		api.PATCH("/tasks/:id/status", taskH.PatchStatus)
		api.DELETE("/tasks/:id", taskH.Delete)

		// Sub-tasks
		api.POST("/tasks/:id/subtasks", subtaskH.Create)
		api.PUT("/subtasks/:id", subtaskH.Update)
		api.DELETE("/subtasks/:id", subtaskH.Delete)

		// Tags (global)
		api.GET("/tags", tagH.List)
		api.POST("/tags", tagH.Create)
		api.PUT("/tags/:id", tagH.Update)
		api.DELETE("/tags/:id", tagH.Delete)

		// Task ↔ Tag associations
		api.POST("/tasks/:id/tags/:tagId", tagH.AssignTag)
		api.DELETE("/tasks/:id/tags/:tagId", tagH.RemoveTag)
	}

	return r
}
