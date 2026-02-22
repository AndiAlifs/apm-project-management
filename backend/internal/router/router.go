package router

import (
	"net/http"

	"apm/backend/internal/config"

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

	// ── API v1 Group (populated in future sprints) ────────────────────────────
	// api := r.Group("/api")
	// Sprint 02: project, task, status, tag handlers registered here.

	return r
}
