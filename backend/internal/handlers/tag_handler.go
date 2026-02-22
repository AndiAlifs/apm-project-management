package handlers

import (
	"net/http"
	"strconv"

	"apm/backend/internal/models"
	"apm/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// TagHandler holds the tag service dependency.
type TagHandler struct {
	svc *services.TagService
}

// NewTagHandler creates a new TagHandler.
func NewTagHandler(svc *services.TagService) *TagHandler {
	return &TagHandler{svc: svc}
}

// List handles GET /api/tags
func (h *TagHandler) List(c *gin.Context) {
	tags, err := h.svc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

// Create handles POST /api/tags
func (h *TagHandler) Create(c *gin.Context) {
	var body models.Tag
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Create(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, body)
}

// Update handles PUT /api/tags/:id
func (h *TagHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var body models.Tag
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

// Delete handles DELETE /api/tags/:id
func (h *TagHandler) Delete(c *gin.Context) {
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

// AssignTag handles POST /api/tasks/:id/tags/:tagId
func (h *TagHandler) AssignTag(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}
	tagID, err := strconv.ParseUint(c.Param("tagId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag id"})
		return
	}
	if err := h.svc.AssignToTask(uint(taskID), uint(tagID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// RemoveTag handles DELETE /api/tasks/:id/tags/:tagId
func (h *TagHandler) RemoveTag(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}
	tagID, err := strconv.ParseUint(c.Param("tagId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag id"})
		return
	}
	if err := h.svc.RemoveFromTask(uint(taskID), uint(tagID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
