# Sprint 05 — Kanban View

**Tag:** `sprint/05-kanban-view`
**Phase:** 3 — Views
**Status:** 🔲 Planned
**Depends on:** Sprint 04 (task modal reusable)

---

## Goals

Build the Kanban board view with drag-and-drop card movement and full column management (add, rename, reorder, delete). This is the centrepiece of the app.

---

## Tasks

**Kanban Board**
- [ ] `KanbanComponent` — horizontal scrollable board layout
- [ ] `KanbanColumnComponent` — individual column with header and card list
- [ ] Fetch project statuses → render one column per status
- [ ] Fetch project tasks → group cards by status

**Task Cards**
- [ ] `TaskCardComponent` — shows: title, priority badge, due date, tag chips, sub-task progress (`2/5`)
- [ ] Click card → open `TaskModalComponent` (reused from Sprint 04)

**Drag and Drop**
- [ ] Integrate Angular CDK `DragDropModule`
- [ ] Drag card within a column to reorder
- [ ] Drag card between columns → call `PATCH /api/tasks/:id/status` to update status
- [ ] Visual feedback: placeholder while dragging, drop zone highlight

**Column Management**
- [ ] Add column button → inline name input → `POST /api/projects/:id/statuses`
- [ ] Rename column → double-click or edit icon → inline edit → `PUT /api/statuses/:id`
- [ ] Reorder columns → drag column header → `PUT /api/statuses/:id` with new order
- [ ] Delete column → icon + confirmation dialog → `DELETE /api/statuses/:id`
- [ ] Guard: cannot delete a column that still has tasks (show error toast)

**Quick Add**
- [ ] "+ Add task" button at bottom of each column → mini inline form
- [ ] Creates task with title, defaulting to that column's status

---

## Features Implemented

- Full Kanban board view with columns from backend
- Task cards with all key metadata visible
- Drag-and-drop between columns (status update persisted)
- Drag-and-drop within column (reorder)
- Column add, rename, reorder, delete
- Quick-add task from within a column
- Reuses task modal from Sprint 04

---

## Completion Criteria

- [ ] Board renders all columns and tasks for the selected project
- [ ] Dragging a card to a new column updates its status in the DB (verify via API or DB)
- [ ] Card count badge in column header is accurate
- [ ] Adding a new column appears immediately without refresh
- [ ] Renaming a column persists after page reload
- [ ] Deleting an empty column removes it; deleting a column with tasks shows an error
- [ ] Quick-add creates task in the correct column
- [ ] Sub-task progress `n/total` shows correctly on card
