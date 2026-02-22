# Sprint 03 — Frontend Shell

**Tag:** `sprint/03-frontend-shell`
**Phase:** 2 — Frontend Foundation
**Status:** 🔲 Planned
**Depends on:** Sprint 01 (backend running)

---

## Goals

Scaffold the Angular application with Tailwind CSS, establish the core design system, and build the app shell (sidebar + top bar + routing). No feature views yet — just the skeleton the rest of the app will live inside.

---

## Tasks

**Scaffold**
- [ ] Create Angular project with Angular CLI (`ng new`)
- [ ] Install and configure Tailwind CSS
- [ ] Configure `environment.ts` with `apiUrl`
- [ ] Set up `HttpClient` in `app.config.ts`
- [ ] Define app routes in `app.routes.ts`

**Design System**
- [ ] Set up Tailwind custom theme: brand colors, font (Google Fonts — Inter), spacing tokens
- [ ] Global `styles.css` with base resets and Tailwind directives
- [ ] Define reusable CSS utility classes for cards, badges, buttons

**Core Services**
- [ ] `ProjectService` — `getProjects()`, `createProject()`, `updateProject()`, `deleteProject()`
- [ ] `TaskService` — `getTasks()`, `getTaskById()`, `createTask()`, `updateTask()`, `deleteTask()`, `moveTaskStatus()`
- [ ] `TagService` — `getTags()`, `createTag()`, `updateTag()`, `deleteTag()`

**App Shell**
- [ ] `SidebarComponent` — project list, active project highlight, create project button
- [ ] `TopBarComponent` — view switcher tabs (Kanban / Backlog / Calendar / List), project name
- [ ] `AppComponent` — sidebar + top bar + `<router-outlet>`
- [ ] Route stubs for: `/kanban`, `/backlog`, `/calendar`, `/list` (empty components for now)

---

## Features Implemented

- Angular app running on `http://localhost:4200`
- Tailwind CSS with custom design tokens
- Sidebar showing all projects (fetched from backend)
- Top bar with view switcher tabs
- Angular routing between four view stubs
- Core API services wired up and injectable

---

## Completion Criteria

- [ ] `ng serve` starts without errors
- [ ] Sidebar fetches and lists projects from the real backend API
- [ ] Clicking a project in the sidebar makes it the active project (highlighted)
- [ ] Clicking view tabs in the top bar navigates to the correct route
- [ ] App looks polished: correct font, color palette, spacing — not a default Angular template
- [ ] No console errors on load
