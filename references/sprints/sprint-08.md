# Sprint 08 — Polish & MVP Complete

**Tag:** `sprint/08-polish`
**Phase:** 4 — Polish
**Status:** 🔲 Planned
**Depends on:** Sprints 01–07 (all views complete)

---

## Goals

Transform the working app into a production-quality MVP. No new features — focus entirely on UX quality, error handling, empty states, animations, and developer documentation.

---

## Tasks

**Loading States**
- [ ] Skeleton loaders for: Backlog table, Kanban board, Calendar grid, List view
- [ ] Loading spinner on task modal save/delete actions
- [ ] Disable action buttons while a request is in-flight

**Error States**
- [ ] Toast notification system (`ToastService`) — success (green), error (red), info (blue)
- [ ] Show toast on: API errors, successful save, successful delete
- [ ] Fallback error message if the entire view fails to load ("Something went wrong, please retry")
- [ ] Network offline detection with banner: "Backend not reachable"

**Empty States**
- [ ] No projects: full-page welcome state with "Create your first project" CTA
- [ ] No tasks in project: illustrated empty state per view (Kanban, Backlog, Calendar, List)
- [ ] No tasks matching filters: "No tasks match your filters" with clear filters button

**Animations & Micro-interactions**
- [ ] Task modal slide-in/out animation
- [ ] Kanban card drop animation (smooth settle)
- [ ] Sidebar project item hover state + active state
- [ ] Priority badge and tag chip hover tooltips
- [ ] Confirmation dialog fade-in

**Accessibility**
- [ ] All interactive elements keyboard-navigable
- [ ] Meaningful `aria-label` on icon buttons
- [ ] Focus trap inside modal dialogs

**Documentation**
- [ ] `README.md` with: project overview, prerequisites, setup steps, running locally, folder structure
- [ ] `.env.example` finalised and documented
- [ ] Brief comment headers in key Go files (`handlers`, `services`, `models`)

---

## Features Implemented

- Toast notification system (success / error / info)
- Skeleton loaders on all views
- Full empty states for all views and filter results
- Smooth animations on modal, Kanban, and sidebar
- Keyboard accessibility + focus management
- Offline detection banner
- Complete README for setup

---

## Completion Criteria

- [ ] App shows skeleton loader while data is fetching (not a blank white page)
- [ ] API error shows a red toast with a readable message
- [ ] New user (no projects) sees the welcome empty state, not a blank sidebar
- [ ] Task modal animates in and out smoothly
- [ ] All buttons can be reached and activated with keyboard (Tab + Enter)
- [ ] README allows a fresh developer to set up and run the app from scratch without asking questions
- [ ] No console errors or warnings in a normal usage session
- [ ] All MVP user stories from the PRD are achievable end-to-end ✅
