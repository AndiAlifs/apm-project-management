# Sprint 04 — Backlog View

**Tag:** `sprint/04-backlog-view`
**Phase:** 2 — Frontend Foundation
**Status:** 🔲 Planned
**Depends on:** Sprint 02 (full API), Sprint 03 (app shell)

---

## Goals

Build the Backlog view — a filterable, sortable table of all tasks — and the Task modal/drawer for creating and editing tasks. This is the first end-to-end feature: data flows from MySQL → Golang → Angular → user.

---

## Tasks

**Backlog View**
- [ ] `BacklogComponent` — table layout with columns: Project, Task Name, Status, Priority, Due Date, Tags
- [ ] Fetch all tasks via `TaskService.getTasks()` with active filters
- [ ] Filter bar: by project, status, priority, tag, due date range
- [ ] Sort by any column (toggle asc/desc)
- [ ] Inline status edit (dropdown in table row)
- [ ] Inline priority edit (dropdown in table row)
- [ ] Pagination or virtual scroll for large task lists
- [ ] Empty state: "No tasks yet — create your first one"

**Task Modal / Drawer**
- [ ] `TaskModalComponent` — right-side drawer (slide-in animation)
- [ ] Fields: title, description (Markdown textarea), project, status, priority, due date, tags, sub-tasks
- [ ] Sub-task checklist: add, toggle done, delete, reorder
- [ ] Tag multi-select with colour chips
- [ ] Save (create / update) and Delete actions
- [ ] Confirmation dialog for delete

**Shared Components**
- [ ] `PriorityBadgeComponent` — coloured badge (Low / Medium / High / Urgent)
- [ ] `TagChipComponent` — coloured pill with tag name
- [ ] `ConfirmDialogComponent` — reusable "are you sure?" modal

---

## Features Implemented

- Backlog list view with live data from backend
- Multi-field filter and sort
- Inline status and priority editing
- Full task create / edit / delete via modal drawer
- Sub-tasks checklist inside task modal
- Tag assignment inside task modal
- Priority and tag chip shared components

---

## Completion Criteria

- [ ] Backlog loads all tasks from the real API on page load
- [ ] Filtering by project + tag + priority returns the correct subset
- [ ] Sorting by Due Date ascending/descending works correctly
- [ ] Creating a task via the modal persists to DB and appears in the list without page reload
- [ ] Editing a task updates the row immediately (optimistic or after confirmation)
- [ ] Deleting a task shows a confirmation dialog, then removes from list
- [ ] Sub-task toggle updates the `is_done` field in the DB
- [ ] Empty state appears when no tasks match current filters
