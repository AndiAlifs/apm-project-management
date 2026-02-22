package handlers

import (
	"net/http"
	"strconv"
	"time"

	"apm/backend/internal/models"
	"apm/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// TaskHandler holds the task service dependency.
type TaskHandler struct {
	svc *services.TaskService
}

// NewTaskHandler creates a new TaskHandler.
func NewTaskHandler(svc *services.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

// List handles GET /api/tasks (backlog with optional filters)
// Query params: projectId, statusId, priority, tagId
func (h *TaskHandler) List(c *gin.Context) {
	filter := services.TaskFilter{}

	if v := c.Query("projectId"); v != "" {
		id, err := strconv.ParseUint(v, 10, 32)
		if err == nil {
			uid := uint(id)
			filter.ProjectID = &uid
		}
	}
	if v := c.Query("statusId"); v != "" {
		id, err := strconv.ParseUint(v, 10, 32)
		if err == nil {
			uid := uint(id)
			filter.StatusID = &uid
		}
	}
	if v := c.Query("priority"); v != "" {
		filter.Priority = &v
	}
	if v := c.Query("tagId"); v != "" {
		id, err := strconv.ParseUint(v, 10, 32)
		if err == nil {
			uid := uint(id)
			filter.TagID = &uid
		}
	}
	if v := c.Query("due_from"); v != "" {
		if t, err := time.Parse("2006-01-02", v); err == nil {
			filter.DueFrom = &t
		}
	}
	if v := c.Query("due_to"); v != "" {
		if t, err := time.Parse("2006-01-02", v); err == nil {
			filter.DueTo = &t
		}
	}

	tasks, err := h.svc.List(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// ListByProject handles GET /api/projects/:id/tasks
func (h *TaskHandler) ListByProject(c *gin.Context) {
	projectID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}
	tasks, err := h.svc.ListByProject(uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetByID handles GET /api/tasks/:id
func (h *TaskHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	task, err := h.svc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Create handles POST /api/tasks
func (h *TaskHandler) Create(c *gin.Context) {
	var body models.Task
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Create(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return fully loaded task
	task, err := h.svc.GetByID(body.ID)
	if err != nil {
		c.JSON(http.StatusCreated, body)
		return
	}
	c.JSON(http.StatusCreated, task)
}

// Update handles PUT /api/tasks/:id
func (h *TaskHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.svc.Update(uint(id), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// PatchStatus handles PATCH /api/tasks/:id/status
func (h *TaskHandler) PatchStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var body struct {
		StatusID uint `json:"statusId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := h.svc.PatchStatus(uint(id), body.StatusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Delete handles DELETE /api/tasks/:id
func (h *TaskHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
