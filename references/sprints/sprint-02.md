# Sprint 02 — Backend API

**Tag:** `sprint/02-backend-api`
**Phase:** 1 — Backend Foundation
**Status:** ✅ Done
**Depends on:** Sprint 01

---

## Goals

Implement all REST API endpoints for the core entities: Projects, Statuses, Tasks, Sub-tasks, and Tags. By the end of this sprint the entire data layer is testable via Postman/Bruno — no frontend needed.

---

## Tasks

**Projects**
- [x] `GET    /api/projects` — list all
- [x] `POST   /api/projects` — create
- [x] `PUT    /api/projects/:id` — update
- [x] `DELETE /api/projects/:id` — delete (cascade tasks)

**Statuses (Kanban Columns)**
- [x] `GET    /api/projects/:id/statuses` — list for project
- [x] `POST   /api/projects/:id/statuses` — create
- [x] `PUT    /api/statuses/:id` — update name / color / order
- [x] `DELETE /api/statuses/:id` — delete

**Tasks**
- [x] `GET    /api/projects/:id/tasks` — list tasks for project
- [x] `GET    /api/tasks` — all tasks (backlog) with query filters: `?project=&status=&priority=&tag=&due_from=&due_to=`
- [x] `POST   /api/tasks` — create task
- [x] `PUT    /api/tasks/:id` — update task
- [x] `DELETE /api/tasks/:id` — delete task
- [x] `PATCH  /api/tasks/:id/status` — move to different status

**Sub-tasks**
- [x] `POST   /api/tasks/:id/subtasks` — add sub-task
- [x] `PUT    /api/subtasks/:id` — update (title / is_done / order)
- [x] `DELETE /api/subtasks/:id` — delete

**Tags**
- [x] `GET    /api/tags` — list all tags
- [x] `POST   /api/tags` — create tag
- [x] `PUT    /api/tags/:id` — update name / color
- [x] `DELETE /api/tags/:id` — delete tag
- [x] `POST   /api/tasks/:id/tags/:tagId` — assign tag to task
- [x] `DELETE /api/tasks/:id/tags/:tagId` — remove tag from task

**Seed Data**
- [x] Write a seed script (or manual SQL) to create one sample project with default statuses and sample tasks for testing

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

- [x] All endpoints return correct HTTP status codes (200, 201, 400, 404, 500)
- [x] Creating a project and immediately fetching it returns the same data
- [x] Deleting a project also removes its tasks (cascade verified in DB)
- [x] Task filter endpoint returns correct subset when filtering by project, tag, and priority simultaneously
- [x] Sub-task `is_done` toggle persists correctly
- [x] Tag assignment and removal reflected in task response
- [x] All endpoints tested manually via API client (Postman / Bruno / Swagger UI)
