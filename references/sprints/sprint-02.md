# Sprint 02 — Backend API

**Tag:** `sprint/02-backend-api`
**Phase:** 1 — Backend Foundation
**Status:** 🔲 Planned
**Depends on:** Sprint 01

---

## Goals

Implement all REST API endpoints for the core entities: Projects, Statuses, Tasks, Sub-tasks, and Tags. By the end of this sprint the entire data layer is testable via Postman/Bruno — no frontend needed.

---

## Tasks

**Projects**
- [ ] `GET    /api/projects` — list all
- [ ] `POST   /api/projects` — create
- [ ] `PUT    /api/projects/:id` — update
- [ ] `DELETE /api/projects/:id` — delete (cascade tasks)

**Statuses (Kanban Columns)**
- [ ] `GET    /api/projects/:id/statuses` — list for project
- [ ] `POST   /api/projects/:id/statuses` — create
- [ ] `PUT    /api/statuses/:id` — update name / color / order
- [ ] `DELETE /api/statuses/:id` — delete

**Tasks**
- [ ] `GET    /api/projects/:id/tasks` — list tasks for project
- [ ] `GET    /api/tasks` — all tasks (backlog) with query filters: `?project=&status=&priority=&tag=&due_from=&due_to=`
- [ ] `POST   /api/tasks` — create task
- [ ] `PUT    /api/tasks/:id` — update task
- [ ] `DELETE /api/tasks/:id` — delete task
- [ ] `PATCH  /api/tasks/:id/status` — move to different status

**Sub-tasks**
- [ ] `POST   /api/tasks/:id/subtasks` — add sub-task
- [ ] `PUT    /api/subtasks/:id` — update (title / is_done / order)
- [ ] `DELETE /api/subtasks/:id` — delete

**Tags**
- [ ] `GET    /api/tags` — list all tags
- [ ] `POST   /api/tags` — create tag
- [ ] `PUT    /api/tags/:id` — update name / color
- [ ] `DELETE /api/tags/:id` — delete tag
- [ ] `POST   /api/tasks/:id/tags/:tagId` — assign tag to task
- [ ] `DELETE /api/tasks/:id/tags/:tagId` — remove tag from task

**Seed Data**
- [ ] Write a seed script (or manual SQL) to create one sample project with default statuses and sample tasks for testing

---

## Features Implemented

- Full CRUD for: Projects, Statuses, Tasks, Sub-tasks, Tags
- Task-tag association endpoints
- Backlog query endpoint with multi-field filtering
- Status move endpoint (used by Kanban drag-and-drop later)
- Request validation via Go struct binding
- Consistent JSON error responses `{ "error": "message" }`

---

## Completion Criteria

- [ ] All endpoints return correct HTTP status codes (200, 201, 400, 404, 500)
- [ ] Creating a project and immediately fetching it returns the same data
- [ ] Deleting a project also removes its tasks (cascade verified in DB)
- [ ] Task filter endpoint returns correct subset when filtering by project, tag, and priority simultaneously
- [ ] Sub-task `is_done` toggle persists correctly
- [ ] Tag assignment and removal reflected in task response
- [ ] All endpoints tested manually via API client (Postman / Bruno / Swagger UI)
