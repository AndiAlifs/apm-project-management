# APM — Agile Personal Manager

A lightweight, single-user personal task management web app. Organise tasks across projects using Kanban, Backlog, Calendar, and List views — with no accounts, no cloud, and no overhead.

---

## Tech Stack

| Layer | Technology |
|-------|------------|
| Frontend | Angular 18 + Tailwind CSS |
| Backend | Golang (Gin + GORM) |
| Database | MySQL 8 |

---

## Prerequisites

- [Go](https://go.dev/dl/) 1.22+
- [Node.js](https://nodejs.org/) 20+ & npm
- [MySQL](https://dev.mysql.com/downloads/) 8.x
- [Angular CLI](https://angular.dev/tools/cli) (`npm install -g @angular/cli`)
- [`air`](https://github.com/air-verse/air) for Go live reload (`go install github.com/air-verse/air@latest`)

---

## Project Structure

```
apm-project-management/
├── backend/                  # Golang REST API
│   ├── cmd/server/main.go    # Entry point
│   ├── internal/
│   │   ├── config/           # Env config
│   │   ├── database/         # MySQL + GORM setup
│   │   ├── models/           # GORM models
│   │   ├── handlers/         # HTTP handlers
│   │   ├── services/         # Business logic
│   │   └── router/           # Route definitions
│   ├── .env                  # Local env vars (not committed)
│   ├── .env.example          # Env var reference
│   └── go.mod
├── frontend/                 # Angular app
│   ├── src/app/
│   │   ├── core/             # Services + models
│   │   ├── features/         # Kanban, Backlog, Calendar, List
│   │   └── shared/           # Reusable components
│   ├── tailwind.config.js
│   └── angular.json
├── references/               # Planning docs
│   ├── prd.md                # Product Requirements Document
│   ├── roadmap.md            # Sprint tracker
│   └── sprints/              # Per-sprint task files
└── README.md
```

---

## Setup

### 1. Database

Create a MySQL database:

```sql
CREATE DATABASE apm CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 2. Backend

```bash
cd backend

# Copy and fill in environment variables
cp .env.example .env

# Download dependencies
go mod tidy

# Run with live reload
air

# Or run directly
go run ./cmd/server
```

Backend runs on **http://localhost:8080**

#### `.env` variables

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=apm
SERVER_PORT=8080
CORS_ORIGIN=http://localhost:4200
```

> Tables are auto-migrated on first run — no manual SQL needed.

### 3. Frontend

```bash
cd frontend

# Install dependencies
npm install

# Start dev server
ng serve
```

Frontend runs on **http://localhost:4200**

---

## API

Base URL: `http://localhost:8080/api`

| Resource | Endpoint |
|----------|----------|
| Projects | `/api/projects` |
| Statuses | `/api/projects/:id/statuses` |
| Tasks | `/api/tasks`, `/api/projects/:id/tasks` |
| Sub-tasks | `/api/tasks/:id/subtasks` |
| Tags | `/api/tags` |

> Full API specification: [`references/prd.md` → Section 10](./references/prd.md)

---

## Sprint Progress

See [`references/roadmap.md`](./references/roadmap.md) for the full sprint plan and current status.

| Sprint | Focus | Status |
|--------|-------|--------|
| 01 | Backend core (Go setup, MySQL schema) | 🔲 Planned |
| 02 | Full REST API | 🔲 Planned |
| 03 | Angular shell (Tailwind, sidebar, routing) | 🔲 Planned |
| 04 | Backlog view + task modal | 🔲 Planned |
| 05 | Kanban + drag-and-drop | 🔲 Planned |
| 06 | Calendar + List views | 🔲 Planned |
| 07 | Tags manager + sub-tasks polish | 🔲 Planned |
| 08 | Polish + MVP complete | 🔲 Planned |

---

## Roadmap (Phase 2)

- Focus timer (Pomodoro) + productivity dashboard
- Desktop packaging (Tauri `.exe`)
- Rich text editor for task descriptions
- Recurring tasks
- Dark / light mode
- Data export (CSV / JSON)
