# Sprint 07 — Tags & Sub-tasks UI

**Tag:** `sprint/07-tags-subtasks`
**Phase:** 3 — Views
**Status:** 🔲 Planned
**Depends on:** Sprint 04 (task modal), Sprint 05–06 (views)

---

## Goals

Polish the tag management experience globally and ensure sub-task progress is accurately reflected across all views. Add tag filtering to Kanban and Calendar views.

---

## Tasks

**Tags Management Page / Panel**
- [ ] `TagsManagerComponent` — accessible from sidebar or settings icon
- [ ] List all tags with name and colour swatch
- [ ] Create tag: name + colour picker
- [ ] Edit tag: inline rename + colour picker
- [ ] Delete tag: with confirmation dialog (warn if tag is in use)
- [ ] Show count "used by N tasks" next to each tag

**Tag Filtering Across Views**
- [ ] Global tag filter chip bar (below top bar or in sidebar)
- [ ] Selecting a tag filters tasks in the active view (Kanban / Backlog / Calendar / List)
- [ ] Multiple tags selectable (AND or OR toggle)
- [ ] Clear all filters button

**Sub-tasks Polish**
- [ ] Sub-task `order` drag-to-reorder inside task modal
- [ ] Progress bar in task modal showing `n / total` sub-tasks complete
- [ ] Sub-task count badge on Kanban cards updates in real time after editing
- [ ] Sub-task checklist in List view inline expansion

**Task Modal Enhancements**
- [ ] Tag multi-select with search/filter inside the dropdown
- [ ] Keyboard shortcut: `Enter` adds a new sub-task in the checklist

---

## Features Implemented

- Full tags management UI (CRUD with colour picker)
- Tag usage count
- Cross-view tag filtering (all four views)
- Sub-task drag-to-reorder
- Sub-task progress bar in modal
- Real-time sub-task badge update on Kanban cards
- Tag search in task modal's multi-select dropdown

---

## Completion Criteria

- [ ] Creating a tag and assigning it to a task shows the colour chip on the task card
- [ ] Deleting a tag removes it from all tasks (cascade) and it disappears from filter bar
- [ ] Filtering by two tags shows only tasks that have both (or either — per mode toggle)
- [ ] Dragging sub-tasks to reorder persists the new order after page reload
- [ ] Sub-task progress on Kanban card matches the actual checklist count
- [ ] Colour picker in tag editor saves the chosen hex/hsl colour correctly
