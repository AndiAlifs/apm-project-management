package main

import (
	"log"

	"apm/backend/internal/config"
	"apm/backend/internal/database"
	"apm/backend/internal/router"
)

func main() {
	// 1. Load configuration from .env / environment variables.
	cfg := config.Load()

	// 2. Connect to MySQL and run database migrations.
	database.Connect(cfg)

	// 3. Set up the HTTP router with CORS and routes.
	r := router.Setup(cfg)

	// 4. Start the server.
	addr := ":" + cfg.ServerPort
	log.Printf("Server starting on http://localhost%s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
