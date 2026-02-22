import { Component } from '@angular/core';

@Component({
    selector: 'app-kanban',
    standalone: true,
    template: `
    <div class="flex h-full items-center justify-center flex-col gap-4 text-slate-600">
      <span class="text-5xl">⬛</span>
      <h2 class="text-lg font-medium text-slate-400">Kanban Board</h2>
      <p class="text-sm">Coming in Phase 3</p>
    </div>
  `,
})
export class KanbanComponent { }
