# Sprint 06 — Calendar & List Views

**Tag:** `sprint/06-calendar-list-views`
**Phase:** 3 — Views
**Status:** 🔲 Planned
**Depends on:** Sprint 04 (task modal reusable)

---

## Goals

Implement the Calendar view (month/week/day) and the List view (grouped, expandable rows per project). Both views reuse the task modal from Sprint 04.

---

## Tasks

**Calendar View**
- [ ] `CalendarComponent` — main container with month/week/day toggle
- [ ] `CalendarMonthComponent` — full month grid, tasks shown on their `due_date`
- [ ] `CalendarWeekComponent` — 7-column week grid
- [ ] `CalendarDayComponent` — single day list of tasks
- [ ] Previous / next navigation (month / week / day)
- [ ] Today button — jumps back to current date
- [ ] Today's date highlighted visually
- [ ] Tasks colour-coded by project colour (toggle available for priority colour)
- [ ] Click on a task chip → open `TaskModalComponent`
- [ ] Click on an empty date → open `TaskModalComponent` with `due_date` pre-filled
- [ ] Overflow handling: "+N more" label if too many tasks on one day

**List View**
- [ ] `ListViewComponent` — tasks for the selected project in grouped rows
- [ ] Group by: Status (default) or Priority (toggle)
- [ ] Expand/collapse group sections
- [ ] Expandable task row: show description + sub-task checklist inline on expand
- [ ] Inline title edit (click to edit)
- [ ] Inline status change (dropdown in row)
- [ ] Add task button per group (inherits group's status or priority)

---

## Features Implemented

- Calendar view with month / week / day toggle
- Tasks placed by due date; colour-coded by project
- Click to create task with pre-filled date
- List view grouped by status or priority
- Expandable rows with inline sub-task checklist
- Inline editing for title and status
- Both views reuse the shared task modal

---

## Completion Criteria

- [ ] Calendar month view shows all tasks with due dates in correct grid cell
- [ ] Navigating months fetches correct task data
- [ ] Clicking an empty date opens task modal with that date pre-filled
- [ ] List view shows all tasks for the active project, grouped by status
- [ ] Expanding a task row shows its description and sub-task checklist
- [ ] Toggling a sub-task in the list view persists to DB
- [ ] Week and day views show correct tasks
