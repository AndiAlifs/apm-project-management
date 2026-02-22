# APM — Sprint Roadmap

> **Legend:** 🔲 Planned · 🔄 In Progress · ✅ Done · ⏸ On Hold

---

## Overview

| Sprint | Tag | Phase | Focus | Status | Sprint Doc |
|--------|-----|-------|-------|--------|------------|
| 01 | `sprint/01-backend-core` | Phase 1 | Go project setup, DB schema, base router | ✅ Done | [sprint-01.md](./sprints/sprint-01.md) |
| 02 | `sprint/02-backend-api` | Phase 1 | Full REST API: projects, tasks, statuses, tags | 🔲 Planned | [sprint-02.md](./sprints/sprint-02.md) |
| 03 | `sprint/03-frontend-shell` | Phase 2 | Angular scaffold, Tailwind, sidebar, routing | 🔲 Planned | [sprint-03.md](./sprints/sprint-03.md) |
| 04 | `sprint/04-backlog-view` | Phase 2 | Backlog list view + task modal (data flows E2E) | 🔲 Planned | [sprint-04.md](./sprints/sprint-04.md) |
| 05 | `sprint/05-kanban-view` | Phase 3 | Kanban board + drag-and-drop + column management | 🔲 Planned | [sprint-05.md](./sprints/sprint-05.md) |
| 06 | `sprint/06-calendar-list-views` | Phase 3 | Calendar view + List view | 🔲 Planned | [sprint-06.md](./sprints/sprint-06.md) |
| 07 | `sprint/07-tags-subtasks` | Phase 3 | Tags UI + sub-tasks checklist + tag filtering | 🔲 Planned | [sprint-07.md](./sprints/sprint-07.md) |
| 08 | `sprint/08-polish` | Phase 4 | Loading states, empty states, animations, README | 🔲 Planned | [sprint-08.md](./sprints/sprint-08.md) |

---

## Phases

### Phase 1 — Backend Foundation
Sprints 01–02. Goal: fully functional REST API with MySQL. Frontend can be tested via API client (Postman/Bruno/Swagger).

### Phase 2 — Frontend Foundation
Sprints 03–04. Goal: Angular app shell + Backlog view with live data from backend.

### Phase 3 — Views
Sprints 05–07. Goal: all four views working end-to-end with tags and sub-tasks.

### Phase 4 — Polish
Sprint 08. Goal: production-quality MVP with no rough edges.

---

## Conventions

- **Branch naming:** `sprint/XX-short-name` (matches sprint tag)
- **Commit prefix:** `[S01]`, `[S02]`, etc. for traceability
- **Done definition:** all completion criteria in the sprint doc are met and manually verified
- **Sprint doc structure:** Goals → Tasks → Features Implemented → Completion Criteria
