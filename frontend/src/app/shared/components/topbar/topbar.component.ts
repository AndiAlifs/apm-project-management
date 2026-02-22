import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';

interface ViewTab {
    label: string;
    route: string;
    icon: string;
}

@Component({
    selector: 'app-topbar',
    standalone: true,
    imports: [RouterLink, RouterLinkActive],
    template: `
    <header class="flex items-center h-14 px-6 border-b border-slate-800 bg-surface-raised shrink-0 gap-6">
      <!-- View Switcher Tabs -->
      <nav class="flex items-center gap-1" role="tablist" aria-label="View switcher">
        @for (tab of tabs; track tab.route) {
          <a
            [routerLink]="tab.route"
            routerLinkActive="bg-primary-600/20 text-primary-300 border-primary-500"
            [routerLinkActiveOptions]="{ exact: false }"
            class="flex items-center gap-2 px-3 py-1.5 rounded-lg text-sm font-medium border border-transparent
                   text-slate-400 hover:text-slate-200 hover:bg-surface-overlay transition-colors duration-150"
            role="tab"
          >
            <span>{{ tab.icon }}</span>
            {{ tab.label }}
          </a>
        }
      </nav>

      <div class="flex-1"></div>

      <!-- Global Actions -->
      <div class="flex items-center gap-2 text-slate-500 text-xs">
        <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
        API Connected
      </div>
    </header>
  `,
})
export class TopbarComponent {
    tabs: ViewTab[] = [
        { label: 'Kanban', route: '/kanban', icon: '⬛' },
        { label: 'Backlog', route: '/backlog', icon: '☰' },
        { label: 'Calendar', route: '/calendar', icon: '📅' },
        { label: 'List', route: '/list', icon: '≡' },
    ];
}
