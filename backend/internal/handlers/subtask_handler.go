package handlers

import (
	"net/http"
	"strconv"

	"apm/backend/internal/models"
	"apm/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// SubTaskHandler holds the subtask service dependency.
type SubTaskHandler struct {
	svc *services.SubTaskService
}

// NewSubTaskHandler creates a new SubTaskHandler.
func NewSubTaskHandler(svc *services.SubTaskService) *SubTaskHandler {
	return &SubTaskHandler{svc: svc}
}

// Create handles POST /api/tasks/:id/subtasks
func (h *SubTaskHandler) Create(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}
	var body models.SubTask
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	body.TaskID = uint(taskID)
	if err := h.svc.Create(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, body)
}

// Update handles PUT /api/subtasks/:id
func (h *SubTaskHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var body models.SubTask
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.svc.Update(uint(id), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete handles DELETE /api/subtasks/:id
func (h *SubTaskHandler) Delete(c *gin.Context) {
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
