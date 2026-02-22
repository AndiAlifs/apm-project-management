import { Routes } from '@angular/router';

export const routes: Routes = [
    { path: '', redirectTo: 'backlog', pathMatch: 'full' },
    {
        path: 'backlog',
        loadComponent: () =>
            import('./features/backlog/backlog.component').then(m => m.BacklogComponent),
    },
    {
        path: 'kanban',
        loadComponent: () =>
            import('./features/kanban/kanban.component').then(m => m.KanbanComponent),
    },
    {
        path: 'calendar',
        loadComponent: () =>
            import('./features/calendar/calendar.component').then(m => m.CalendarComponent),
    },
    {
        path: 'list',
        loadComponent: () =>
            import('./features/list/list.component').then(m => m.ListComponent),
    },
    { path: '**', redirectTo: 'backlog' },
];
