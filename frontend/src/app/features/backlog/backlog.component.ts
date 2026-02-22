import {
    Component, OnInit, signal, inject, computed
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Task, Project, Status, Tag, TaskFilter, Priority } from '../../core/models/models';
import { TaskService } from '../../core/services/task.service';
import { ProjectService } from '../../core/services/project.service';
import { TagService } from '../../core/services/tag.service';
import { TaskModalComponent } from '../../shared/components/task-modal/task-modal.component';

type SortKey = 'title' | 'priority' | 'dueDate' | 'project' | 'status';

const PRIORITY_RANK: Record<Priority, number> = { low: 0, medium: 1, high: 2, urgent: 3 };

@Component({
    selector: 'app-backlog',
    standalone: true,
    imports: [CommonModule, FormsModule, TaskModalComponent],
    template: `
    <div class="h-full flex flex-col p-6 gap-4">
      <!-- Header -->
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-xl font-semibold text-slate-100">Backlog</h1>
          <p class="text-xs text-slate-500 mt-0.5">{{ filteredTasks().length }} tasks</p>
        </div>
        <button class="btn-primary" (click)="openCreateModal()" id="new-task-btn">
          + New Task
        </button>
      </div>

      <!-- Filter Bar -->
      <div class="flex flex-wrap gap-3 items-center">
        <select id="filter-project" class="select-base w-40" [(ngModel)]="filter.projectId" (change)="applyFilter()">
          <option [ngValue]="undefined">All Projects</option>
          @for (p of projects(); track p.id) {
            <option [ngValue]="p.id">{{ p.name }}</option>
          }
        </select>
        <select id="filter-priority" class="select-base w-36" [(ngModel)]="filter.priority" (change)="applyFilter()">
          <option [ngValue]="undefined">All Priorities</option>
          <option value="urgent">🔴 Urgent</option>
          <option value="high">🟠 High</option>
          <option value="medium">🔵 Medium</option>
          <option value="low">⚪ Low</option>
        </select>
        <select id="filter-tag" class="select-base w-36" [(ngModel)]="filter.tagId" (change)="applyFilter()">
          <option [ngValue]="undefined">All Tags</option>
          @for (t of tags(); track t.id) {
            <option [ngValue]="t.id">{{ t.name }}</option>
          }
        </select>
        @if (filter.projectId || filter.priority || filter.tagId) {
          <button class="btn-ghost text-xs" (click)="clearFilter()">✕ Clear</button>
        }
      </div>

      <!-- Table -->
      <div class="flex-1 overflow-auto rounded-xl border border-slate-800">
        @if (loading()) {
          <div class="flex items-center justify-center h-48 text-slate-500">Loading tasks…</div>
        } @else if (filteredTasks().length === 0) {
          <div class="flex flex-col items-center justify-center h-48 gap-3 text-slate-600">
            <span class="text-4xl">📭</span>
            <p class="text-sm">No tasks yet. Click <strong class="text-slate-400">+ New Task</strong> to get started.</p>
          </div>
        } @else {
          <table class="w-full text-sm border-collapse">
            <thead class="sticky top-0 bg-surface-raised border-b border-slate-800">
              <tr>
                @for (col of columns; track col.key) {
                  <th
                    class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider cursor-pointer hover:text-slate-300 select-none"
                    (click)="sortBy(col.key)"
                  >
                    {{ col.label }}
                    @if (sortKey() === col.key) {
                      <span class="ml-1">{{ sortDir() === 'asc' ? '↑' : '↓' }}</span>
                    }
                  </th>
                }
                <th class="px-4 py-3 w-10"></th>
              </tr>
            </thead>
            <tbody>
              @for (task of filteredTasks(); track task.id) {
                <tr
                  class="border-b border-slate-800/60 hover:bg-surface-raised/60 cursor-pointer transition-colors"
                  (click)="openEditModal(task)"
                >
                  <!-- Project -->
                  <td class="px-4 py-3">
                    <div class="flex items-center gap-2">
                      <span class="w-2 h-2 rounded-full shrink-0"
                            [style.background-color]="task.project?.color || '#94A3B8'"></span>
                      <span class="text-slate-400 text-xs truncate max-w-24">{{ task.project?.name || '—' }}</span>
                    </div>
                  </td>
                  <!-- Title -->
                  <td class="px-4 py-3 font-medium text-slate-100 max-w-xs">
                    <div class="truncate">{{ task.title }}</div>
                    @if (task.subTasks && task.subTasks.length > 0) {
                      <div class="text-xs text-slate-500 mt-0.5">
                        {{ doneCount(task) }}/{{ task.subTasks!.length }} sub-tasks
                      </div>
                    }
                  </td>
                  <!-- Status -->
                  <td class="px-4 py-3">
                    <span class="px-2 py-0.5 rounded-full text-xs font-medium"
                          [style.background-color]="(task.status?.color || '#94A3B8') + '33'"
                          [style.color]="task.status?.color || '#94A3B8'">
                      {{ task.status?.name || '—' }}
                    </span>
                  </td>
                  <!-- Priority -->
                  <td class="px-4 py-3">
                    <span class="px-2 py-0.5 rounded-full text-xs font-medium"
                          [ngClass]="'badge-priority-' + task.priority">
                      {{ task.priority | titlecase }}
                    </span>
                  </td>
                  <!-- Due Date -->
                  <td class="px-4 py-3 text-xs text-slate-400">
                    {{ task.dueDate ? formatDate(task.dueDate) : '—' }}
                  </td>
                  <!-- Tags -->
                  <td class="px-4 py-3">
                    <div class="flex flex-wrap gap-1">
                      @for (tag of task.tags?.slice(0, 3); track tag.id) {
                        <span class="px-1.5 py-0.5 rounded text-xs font-medium"
                              [style.background-color]="tag.color + '33'"
                              [style.color]="tag.color">
                          {{ tag.name }}
                        </span>
                      }
                      @if ((task.tags?.length || 0) > 3) {
                        <span class="text-xs text-slate-500">+{{ (task.tags?.length || 0) - 3 }}</span>
                      }
                    </div>
                  </td>
                  <!-- Actions -->
                  <td class="px-4 py-3 text-right" (click)="$event.stopPropagation()">
                    <button class="text-slate-600 hover:text-red-400 transition-colors text-xs px-2"
                            (click)="deleteTask(task)" title="Delete task">✕</button>
                  </td>
                </tr>
              }
            </tbody>
          </table>
        }
      </div>
    </div>

    <!-- Task Modal/Drawer -->
    @if (showModal()) {
      <app-task-modal
        [task]="selectedTask()"
        [projects]="projects()"
        [tags]="tags()"
        (saved)="onTaskSaved($event)"
        (closed)="showModal.set(false)"
      />
    }
  `,
})
export class BacklogComponent implements OnInit {
    private taskSvc = inject(TaskService);
    private projectSvc = inject(ProjectService);
    private tagSvc = inject(TagService);

    tasks = signal<Task[]>([]);
    projects = signal<Project[]>([]);
    tags = signal<Tag[]>([]);
    loading = signal(true);

    showModal = signal(false);
    selectedTask = signal<Task | null>(null);

    filter: TaskFilter = {};
    sortKey = signal<SortKey>('dueDate');
    sortDir = signal<'asc' | 'desc'>('asc');

    columns: { key: SortKey; label: string }[] = [
        { key: 'project', label: 'Project' },
        { key: 'title', label: 'Task' },
        { key: 'status', label: 'Status' },
        { key: 'priority', label: 'Priority' },
        { key: 'dueDate', label: 'Due Date' },
    ];

    filteredTasks = computed(() => {
        const key = this.sortKey();
        const dir = this.sortDir();
        return [...this.tasks()].sort((a, b) => {
            let av: string | number;
            let bv: string | number;
            if (key === 'project') { av = a.project?.name || ''; bv = b.project?.name || ''; }
            else if (key === 'status') { av = a.status?.name || ''; bv = b.status?.name || ''; }
            else if (key === 'priority') { av = PRIORITY_RANK[a.priority]; bv = PRIORITY_RANK[b.priority]; }
            else if (key === 'dueDate') { av = a.dueDate || '9999'; bv = b.dueDate || '9999'; }
            else { av = (a as unknown as Record<string, string>)[key] || ''; bv = (b as unknown as Record<string, string>)[key] || ''; }
            if (av < bv) return dir === 'asc' ? -1 : 1;
            if (av > bv) return dir === 'asc' ? 1 : -1;
            return 0;
        });
    });

    ngOnInit(): void {
        this.loadData();
    }

    loadData(): void {
        this.loading.set(true);
        this.projectSvc.list().subscribe(p => this.projects.set(p));
        this.tagSvc.list().subscribe(t => this.tags.set(t));
        this.taskSvc.list(this.filter).subscribe({
            next: t => { this.tasks.set(t); this.loading.set(false); },
            error: () => this.loading.set(false),
        });
    }

    applyFilter(): void { this.loadData(); }

    clearFilter(): void {
        this.filter = {};
        this.loadData();
    }

    sortBy(key: SortKey): void {
        if (this.sortKey() === key) {
            this.sortDir.update(d => d === 'asc' ? 'desc' : 'asc');
        } else {
            this.sortKey.set(key);
            this.sortDir.set('asc');
        }
    }

    openCreateModal(): void {
        this.selectedTask.set(null);
        this.showModal.set(true);
    }

    openEditModal(task: Task): void {
        this.selectedTask.set(task);
        this.showModal.set(true);
    }

    onTaskSaved(task: Task): void {
        this.showModal.set(false);
        this.loadData();
    }

    deleteTask(task: Task): void {
        if (!confirm(`Delete "${task.title}"?`)) return;
        this.taskSvc.delete(task.id).subscribe(() => this.loadData());
    }

    doneCount(task: Task): number {
        return task.subTasks?.filter(s => s.isDone).length ?? 0;
    }

    formatDate(d: string): string {
        return new Date(d).toLocaleDateString('en-GB', { day: 'numeric', month: 'short', year: 'numeric' });
    }
}
