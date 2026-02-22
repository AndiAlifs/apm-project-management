# Sprint 01 — Backend Core

**Tag:** `sprint/01-backend-core`
**Phase:** 1 — Backend Foundation
**Status:** 🔲 Planned

---

## Goals

Stand up the Go backend with a working MySQL connection, full database schema, and a running HTTP server with health check. No business logic yet — just solid foundations.

---

## Tasks

- [ ] Initialise Go module (`go mod init`)
- [ ] Add dependencies: `gin` (or `chi`), `gorm`, `gorm/driver/mysql`, `godotenv`
- [ ] Create project directory structure as per PRD
- [ ] Implement `config.go` — load `.env` variables
- [ ] Implement `db.go` — MySQL connection + GORM auto-migrate
- [ ] Define all GORM models: `Project`, `Status`, `Task`, `SubTask`, `Tag`, `TaskTag`
- [ ] Implement `router.go` — base router with CORS middleware
- [ ] Add `GET /health` endpoint returning `{ "status": "ok" }`
- [ ] Create `.env.example` with all required variables
- [ ] Verify `air` live-reload works (`air` config file)
- [ ] Confirm all tables are created in MySQL on first run

---

## Features Implemented

- Go project fully scaffolded
- MySQL connection established via GORM
- All database tables auto-migrated on startup
- HTTP server running on `PORT` from `.env`
- CORS configured for `http://localhost:4200`
- Health check endpoint: `GET /health`

---

## Completion Criteria

- [ ] Running `go run ./cmd/server` (or `air`) starts the server without errors
- [ ] All 6 tables exist in MySQL after first run: `projects`, `statuses`, `tasks`, `sub_tasks`, `tags`, `task_tags`
- [ ] `GET http://localhost:8080/health` returns `200 OK` with `{ "status": "ok" }`
- [ ] `.env.example` documents all required environment variables
