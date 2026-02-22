# APM — Personal Task Management Tool
# Product Requirements Document

## 1. Executive Summary

APM (Agile Personal Manager) is a single-user, local-first personal task management web application inspired by Notion and Trello. It lets users organise tasks across multiple projects using four interchangeable views — Kanban, Backlog, Calendar, and List — with rich metadata per task (priority, due date, tags, sub-tasks) and optional productivity tracking.

The core value proposition is **focused simplicity**: a lightweight, native-feeling tool that runs in the browser and can be packaged as a desktop `.exe`, with no accounts, no cloud, and no unnecessary complexity.

**MVP Goal:** Deliver a functional task management application with multi-project support, four task views, customisable Kanban columns, fully customisable tags, and a clean Angular + Tailwind frontend backed by a Golang REST API and MySQL database.

---

## 2. Mission

**Mission Statement:** Provide a fast, distraction-free personal workspace to capture, organise, and track tasks across all areas of life.

### Core Principles

1. **Simplicity First** — Every feature must earn its place. Launch fast, iterate deliberately.
2. **Views as First-Class Citizens** — Switching between Kanban, Backlog, Calendar, and List must feel instant and natural.
3. **Owned Data** — All data lives locally on the user's machine; no cloud dependency.
4. **Native Feel in a Browser** — SPA routing, smooth animations, and desktop packaging make it feel like a real app.
5. **Extensibility** — Columns, tags, and statuses are customisable; the tool adapts to the user's workflow.

---

## 3. Target Users

### Primary Persona: The Productivity-Conscious Individual

- **Who:** A single user running the app locally for personal and professional task management
- **Technical Comfort:** Comfortable setting up a local dev environment or running an `.exe`
- **Goals:**
  - Track tasks across multiple contexts: day job, freelance clients, and personal life
  - Visualise work in different ways depending on the situation (board vs. list vs. calendar)
  - Measure personal productivity over time without switching to a separate tool
- **Pain Points:**
  - Notion is powerful but slow and cloud-dependent
  - Trello lacks the rich task detail they need (sub-tasks, descriptions, time tracking)
  - Existing tools require accounts and subscriptions for basic features
  - Context-switching between work, freelance, and personal tasks is disorganised

---

## 4. MVP Scope

### In Scope

**Core Functionality**
- ✅ Create, edit, and delete Projects
- ✅ Create, edit, and delete Tasks with: title, description, priority, due date, status, tags, sub-tasks
- ✅ Customisable Kanban columns (add, rename, reorder, delete per project)
- ✅ Kanban view with drag-and-drop between columns
- ✅ Backlog list view across all projects with sort and filter
- ✅ Calendar view with tasks placed by due date (month/week/day)
- ✅ List view per project with expandable task rows
- ✅ Fully customisable tags (create, rename, colour, delete)
- ✅ Multi-tag support per task
- ✅ Sub-tasks checklist per task (add, complete, reorder, delete)
- ✅ Task priority levels: Low / Medium / High / Urgent

**Technical**
- ✅ Angular frontend with Tailwind CSS
- ✅ Golang REST API backend
- ✅ MySQL database for persistence
- ✅ RESTful API design
- ✅ Local development setup

### Out of Scope

**Deferred to Phase 2**
- ❌ Productivity dashboard and focus timer
- ❌ Recurring tasks
- ❌ Rich WYSIWYG editor (plain Markdown for MVP)
- ❌ Desktop packaging (Tauri/Electron `.exe`)
- ❌ Global search
- ❌ Dark / light mode toggle
- ❌ Data export (CSV / JSON)
- ❌ Authentication / multi-user
- ❌ Cloud sync
- ❌ Notifications / reminders
- ❌ Mobile app (responsive web only)

---

## 5. User Stories

### Primary User Stories

1. **As a user, I want to create a project, so that I can group related tasks together.**
   - Example: Create a "Freelance — Client X" project with a distinct colour

2. **As a user, I want to add tasks to a project with a title, description, priority, and due date, so that I have all relevant context in one place.**
   - Example: Add "Design landing page" with High priority, due next Friday

3. **As a user, I want to see my tasks on a Kanban board, so that I can see the overall progress at a glance.**
   - Example: Columns "Backlog → In Progress → Review → Done" with draggable cards

4. **As a user, I want to customise Kanban columns per project, so that the board reflects my actual workflow.**
   - Example: Add a "Blocked" column between "In Progress" and "Review"

5. **As a user, I want to see all tasks in a flat backlog list, so that I can triage and prioritise across projects.**
   - Example: Filter by tag "freelance" and sort by due date ascending

6. **As a user, I want to view tasks on a calendar, so that I can understand my workload by date.**
   - Example: See three tasks due on Thursday highlighted on the weekly view

7. **As a user, I want to assign custom tags to tasks, so that I can cross-cut tasks by context.**
   - Example: Tag a task with both "job" and "urgent"

8. **As a user, I want to add sub-tasks to a task, so that I can break work into manageable steps.**
   - Example: "Launch campaign" with sub-tasks: write copy, design banner, schedule post

---

## 6. Core Architecture & Patterns

### High-Level Architecture

```
┌─────────────────────┐     HTTP/JSON      ┌─────────────────────┐
│                     │ ◄────────────────► │                     │
│  Angular + Tailwind │                    │   Golang (REST API)  │
│     (Frontend)      │                    │      (Backend)       │
│     Port 4200       │                    │      Port 8080       │
└─────────────────────┘                    └──────────┬──────────┘
                                                      │
                                                      ▼
                                           ┌─────────────────────┐
                                           │        MySQL         │
                                           │      (Database)      │
                                           │      Port 3306       │
                                           └─────────────────────┘
```

### Directory Structure

```
apm-project-management/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go              # Entry point
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go            # Env config loader
│   │   ├── database/
│   │   │   └── db.go                # MySQL connection & migrations
│   │   ├── models/
│   │   │   ├── project.go
│   │   │   ├── task.go
│   │   │   ├── subtask.go
│   │   │   ├── status.go
│   │   │   └── tag.go
│   │   ├── handlers/
│   │   │   ├── project_handler.go
│   │   │   ├── task_handler.go
│   │   │   ├── status_handler.go
│   │   │   └── tag_handler.go
│   │   ├── services/
│   │   │   ├── project_service.go
│   │   │   └── task_service.go
│   │   └── router/
│   │       └── router.go            # Route definitions
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   │   ├── app/
│   │   │   ├── core/
│   │   │   │   ├── services/        # API client services
│   │   │   │   │   ├── project.service.ts
│   │   │   │   │   ├── task.service.ts
│   │   │   │   │   └── tag.service.ts
│   │   │   │   └── models/          # TypeScript interfaces
│   │   │   ├── features/
│   │   │   │   ├── kanban/          # Kanban board view
│   │   │   │   ├── backlog/         # Backlog list view
│   │   │   │   ├── calendar/        # Calendar view
│   │   │   │   └── list/            # List view
│   │   │   ├── shared/
│   │   │   │   └── components/      # Reusable UI components
│   │   │   │       ├── task-card/
│   │   │   │       ├── task-modal/
│   │   │   │       ├── tag-chip/
│   │   │   │       └── sidebar/
│   │   │   ├── app.component.ts
│   │   │   ├── app.routes.ts
│   │   │   └── app.config.ts
│   │   ├── styles.css               # Tailwind base imports
│   │   └── main.ts
│   ├── angular.json
│   ├── package.json
│   ├── tailwind.config.js
│   └── tsconfig.json
│
├── references/
│   └── prd.md                       # This document
├── .gitignore
└── README.md
```

### Key Design Patterns

- **Feature Module Architecture** — Angular features are self-contained (kanban, backlog, calendar, list)
- **Service Layer** — All API calls centralised in `core/services/`
- **Repository Pattern (Backend)** — Database access abstracted behind service layer in Go
- **Optimistic UI** — UI updates immediately; server confirms in background
- **Drag-and-Drop with CDK** — Angular CDK DragDrop for Kanban card movement

---

## 7. Features

### 7.1 Project Management

**Purpose:** Group and organise tasks by project

**Operations:**
- Create project with name, description, colour, and icon
- Edit project metadata
- Delete project (with confirmation; cascades to tasks)
- Switch active project via sidebar

**Key Features:**
- Projects listed in collapsible sidebar
- Each project has its own set of Kanban columns (statuses)
- Project colour used as accent throughout the view

---

### 7.2 Task Management

**Purpose:** Create and manage tasks within projects

**Operations:**
- Create task with: title (required), description (Markdown), project, status, priority, due date, tags, sub-tasks
- Edit all task fields
- Move task between projects
- Delete task (with confirmation)

**Key Features:**
- Task detail opens in a right-side modal/drawer — no page reload
- Sub-task checklist with progress indicator (e.g., `3/5 done`)
- `started_at` auto-set when status first moves from To Do → active
- `completed_at` auto-set when status moves to Done column

---

### 7.3 Kanban View

**Purpose:** Visualise and manage tasks as a board

**Features:**
- Columns = Statuses defined per project
- Default columns per new project: `To Do → In Progress → Done`
- Add / rename / reorder / delete columns inline
- Drag-and-drop cards between columns to update status
- Task cards show: title, priority badge, due date, tag chips, sub-task progress
- Quick-add task button at the bottom of each column
- Column card count badge

---

### 7.4 Backlog View

**Purpose:** Triage and manage all tasks in a flat table

**Features:**
- Table columns: Project, Task Name, Status, Priority, Due Date, Tags
- Filter by: project, status, priority, tag, due date range
- Sort by any column (ascending / descending)
- Inline status and priority editing
- Bulk actions: update status, assign/remove tags, delete
- Pagination or virtual scroll for large lists

---

### 7.5 Calendar View

**Purpose:** Visualise tasks by due date

**Features:**
- Month / Week / Day toggle
- Tasks placed on their `due_date`
- Colour-coded by project or priority (user toggle)
- Click empty date → quick create task with that due date pre-filled
- Click task → open task detail drawer
- Today highlighted

---

### 7.6 List View

**Purpose:** Ordered, detail-rich view of tasks within one project

**Features:**
- Grouped by status or priority
- Expandable rows showing description + sub-task checklist inline
- Inline edit of title and status
- Keyboard-friendly (tab through fields)

---

### 7.7 Tags

**Purpose:** Cross-cutting labels that span projects and statuses

**Features:**
- Create, rename, recolour, and delete tags globally
- Assign multiple tags to a task
- Filter tasks by tag in Backlog and Calendar views
- Tag chips displayed on task cards and rows

---

## 8. Technology Stack

### Backend

| Component | Technology | Notes |
|-----------|------------|-------|
| Language | Go (Golang) | Latest stable |
| HTTP Router | `chi` or `gin` | Lightweight, idiomatic |
| ORM / Query | `GORM` | MySQL driver |
| Database | MySQL | Version 8.x |
| Validation | Go struct tags + custom | Request validation |
| Config | `godotenv` | `.env` file loading |

### Frontend

| Component | Technology | Version |
|-----------|------------|---------|
| Framework | Angular | ^18.x |
| Language | TypeScript | ^5.x |
| Styling | Tailwind CSS | ^3.x |
| Drag & Drop | Angular CDK | bundled |
| HTTP Client | Angular `HttpClient` | bundled |
| State | Angular Signals / Services | built-in |
| Icons | Heroicons / Lucide Angular | latest |
| Date Utilities | `date-fns` | ^3.x |
| Markdown | `ngx-markdown` | latest |

### Development Tools

| Tool | Purpose |
|------|---------|
| `air` (Go) | Live reload for backend |
| Angular CLI | Scaffolding and dev server |
| MySQL Workbench / DBeaver | Database inspection |
| ESLint + Prettier | Frontend code quality |
| `golangci-lint` | Backend linting |

---

## 9. Security & Configuration

### Security Scope

**In Scope:**
- ✅ Input validation on all API endpoints (Go struct tags)
- ✅ SQL injection prevention (GORM parameterised queries)
- ✅ CORS configured for local frontend origin only

**Out of Scope (single-user local tool):**
- ❌ Authentication / authorisation
- ❌ HTTPS (local only)
- ❌ Rate limiting
- ❌ CSRF protection

### Configuration

**Backend `.env`:**
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=secret
DB_NAME=apm
SERVER_PORT=8080
CORS_ORIGIN=http://localhost:4200
```

**Frontend `environment.ts`:**
```typescript
export const environment = {
  production: false,
  apiUrl: 'http://localhost:8080/api'
};
```

### Running Locally

- Backend: `air` (or `go run ./cmd/server`)
- Frontend: `ng serve`

---

## 10. API Specification

### Base URL
```
http://localhost:8080/api
```

### Endpoints

#### Projects

| Method | Path | Description |
|--------|------|-------------|
| GET | `/projects` | List all projects |
| POST | `/projects` | Create a project |
| PUT | `/projects/:id` | Update a project |
| DELETE | `/projects/:id` | Delete a project |

#### Statuses (Kanban Columns)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/projects/:id/statuses` | List statuses for a project |
| POST | `/projects/:id/statuses` | Create a status |
| PUT | `/statuses/:id` | Update / reorder a status |
| DELETE | `/statuses/:id` | Delete a status |

#### Tasks

| Method | Path | Description |
|--------|------|-------------|
| GET | `/projects/:id/tasks` | List tasks for a project |
| GET | `/tasks` | List all tasks (backlog) with filters |
| POST | `/tasks` | Create a task |
| PUT | `/tasks/:id` | Update a task |
| DELETE | `/tasks/:id` | Delete a task |
| PATCH | `/tasks/:id/status` | Move task to a different status |

#### Sub-tasks

| Method | Path | Description |
|--------|------|-------------|
| POST | `/tasks/:id/subtasks` | Add a sub-task |
| PUT | `/subtasks/:id` | Update a sub-task |
| DELETE | `/subtasks/:id` | Delete a sub-task |

#### Tags

| Method | Path | Description |
|--------|------|-------------|
| GET | `/tags` | List all tags |
| POST | `/tags` | Create a tag |
| PUT | `/tags/:id` | Update a tag |
| DELETE | `/tags/:id` | Delete a tag |
| POST | `/tasks/:id/tags/:tagId` | Assign tag to task |
| DELETE | `/tasks/:id/tags/:tagId` | Remove tag from task |

**Example Task Response:**
```json
{
  "id": 1,
  "title": "Design landing page",
  "description": "# Brief\nDesign the new homepage...",
  "project": { "id": 2, "name": "Freelance — Client X", "color": "#6366F1" },
  "status": { "id": 4, "name": "In Progress" },
  "priority": "high",
  "dueDate": "2026-03-01",
  "tags": [{ "id": 1, "name": "freelance", "color": "#F59E0B" }],
  "subTasks": [
    { "id": 1, "title": "Write copy", "isDone": true },
    { "id": 2, "title": "Design banner", "isDone": false }
  ],
  "startedAt": "2026-02-20T09:00:00Z",
  "completedAt": null,
  "createdAt": "2026-02-18T08:00:00Z"
}
```

---

## 11. Success Criteria

### MVP Success Definition

The MVP is successful when a user can:
1. Create a project and customise its Kanban columns
2. Add tasks with full metadata and sub-tasks
3. Drag tasks across Kanban columns to update status
4. Switch to Backlog, Calendar, or List view and see the same data
5. Create and assign custom tags to tasks
6. Filter and sort the backlog by any field
7. All data persists across sessions (MySQL)

### Functional Requirements

- ✅ Full CRUD for projects, tasks, statuses, tags, sub-tasks
- ✅ Kanban drag-and-drop updates task status in real time
- ✅ Backlog filters and sorts correctly by all fields
- ✅ Calendar places tasks on correct due dates
- ✅ Sub-task progress accurately calculated
- ✅ Tags filterable across all views
- ✅ No data loss on normal usage

### Quality Indicators

- Initial page load under 2 seconds (local)
- Drag-and-drop action feedback under 100 ms
- Works in Chrome and Edge (primary targets)
- Responsive layout (desktop primary; tablet acceptable)

---

## 12. Implementation Phases

### Phase 1: Backend Foundation

**Goal:** Fully functional REST API with MySQL persistence

**Deliverables:**
- ✅ Go project scaffolded with `chi`/`gin` router and GORM
- ✅ MySQL schema with all tables and indexes
- ✅ CRUD endpoints: projects, statuses, tasks, sub-tasks, tags
- ✅ CORS configured for local frontend
- ✅ API tested via Swagger / Bruno / Postman

**Validation:** All entities creatable and retrievable via API tool

---

### Phase 2: Frontend Foundation

**Goal:** Angular app shell with data flowing from backend

**Deliverables:**
- ✅ Angular project scaffolded with Tailwind CSS
- ✅ App shell: sidebar (projects) + top bar (view switcher)
- ✅ Core services: `ProjectService`, `TaskService`, `TagService`
- ✅ Backlog list view with filter/sort (simplest view first)
- ✅ Task creation and edit modal/drawer

**Validation:** Can create projects and tasks and see them in the backlog

---

### Phase 3: Views

**Goal:** All four views functional

**Deliverables:**
- ✅ Kanban view with drag-and-drop (Angular CDK)
- ✅ Customisable columns (add, rename, reorder, delete)
- ✅ Calendar view (month/week/day)
- ✅ List view with expandable rows and inline edit
- ✅ Tags UI: create, colour-pick, assign to task, filter by

**Validation:** All user stories achievable end-to-end

---

### Phase 4: Polish

**Goal:** Production-quality MVP

**Deliverables:**
- ✅ Loading skeletons and error states
- ✅ Empty states (no projects, no tasks)
- ✅ Confirmation dialogs for destructive actions
- ✅ Subtle transitions and micro-animations
- ✅ README with full setup instructions

**Validation:** Smooth, delightful UX with no rough edges

---

## 13. Future Considerations

### Post-MVP Enhancements (Phase 2)

- **Productivity Dashboard** — Focus timer (Pomodoro), daily focus summary, velocity chart, task completion time
- **Desktop Packaging** — Tauri `.exe` wrapper for native-feel install
- **Rich Text Editor** — WYSIWYG task descriptions (TipTap or similar)
- **Recurring Tasks** — Daily, weekly, or custom repeat schedules
- **Global Search** — Full-text search across all tasks and projects
- **Dark / Light Mode** — System preference detection with manual toggle
- **Data Export** — CSV / JSON export of projects and tasks
- **Reminders** — Local browser notifications for due dates

### Technical Improvements

- **Database Migrations** — Structured schema versioning with `golang-migrate`
- **PWA Support** — Installable web app with offline capability
- **Integration Tests** — API contract testing
- **CI/CD** — GitHub Actions for lint and test on push

---

## 14. Risks & Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| **Drag-and-drop complexity** | Kanban feels broken if ordering is unreliable | Use Angular CDK DragDrop (battle-tested); store explicit `order` field in DB |
| **MySQL setup friction** | Blocks running the app locally | Document setup clearly in README; provide Docker Compose as alternative |
| **Scope creep (productivity features)** | MVP never ships | Strictly defer Phase 2 items; label them clearly in code as `// TODO: Phase 2` |
| **State management complexity** | Views get out of sync after mutations | Centralise state in Angular services; refresh relevant queries after mutations |
| **Date/timezone issues** | Tasks appear on wrong day in calendar | Normalise all dates to local date strings (`YYYY-MM-DD`); use `date-fns` |

---

## 15. Appendix

### Database Schema

```sql
CREATE TABLE projects (
    id         INT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    description TEXT,
    color      VARCHAR(7) DEFAULT '#6366F1',
    icon       VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE statuses (
    id         INT AUTO_INCREMENT PRIMARY KEY,
    project_id INT NOT NULL,
    name       VARCHAR(100) NOT NULL,
    color      VARCHAR(7) DEFAULT '#94A3B8',
    `order`    INT NOT NULL DEFAULT 0,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

CREATE TABLE tasks (
    id           INT AUTO_INCREMENT PRIMARY KEY,
    project_id   INT NOT NULL,
    status_id    INT NOT NULL,
    title        VARCHAR(500) NOT NULL,
    description  TEXT,
    priority     ENUM('low','medium','high','urgent') DEFAULT 'medium',
    due_date     DATE,
    started_at   TIMESTAMP NULL,
    completed_at TIMESTAMP NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (status_id) REFERENCES statuses(id)
);

CREATE TABLE sub_tasks (
    id      INT AUTO_INCREMENT PRIMARY KEY,
    task_id INT NOT NULL,
    title   VARCHAR(500) NOT NULL,
    is_done BOOLEAN DEFAULT FALSE,
    `order` INT NOT NULL DEFAULT 0,
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
);

CREATE TABLE tags (
    id    INT AUTO_INCREMENT PRIMARY KEY,
    name  VARCHAR(100) NOT NULL UNIQUE,
    color VARCHAR(7) DEFAULT '#F59E0B'
);

CREATE TABLE task_tags (
    task_id INT NOT NULL,
    tag_id  INT NOT NULL,
    PRIMARY KEY (task_id, tag_id),
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

-- Indexes
CREATE INDEX idx_tasks_project ON tasks(project_id);
CREATE INDEX idx_tasks_status  ON tasks(status_id);
CREATE INDEX idx_tasks_due     ON tasks(due_date);
CREATE INDEX idx_statuses_project ON statuses(project_id);
```

### Key Dependencies & References

- [Angular Documentation](https://angular.dev/)
- [Tailwind CSS Documentation](https://tailwindcss.com/)
- [Angular CDK — Drag and Drop](https://material.angular.io/cdk/drag-drop/overview)
- [Gin Web Framework (Go)](https://gin-gonic.com/)
- [GORM — Go ORM](https://gorm.io/)
- [date-fns Documentation](https://date-fns.org/)
- [MySQL 8 Documentation](https://dev.mysql.com/doc/refman/8.0/en/)
