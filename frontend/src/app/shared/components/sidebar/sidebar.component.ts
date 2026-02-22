import { Component, OnInit, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Project } from '../../../core/models/models';
import { ProjectService } from '../../../core/services/project.service';

@Component({
    selector: 'app-sidebar',
    standalone: true,
    imports: [CommonModule, FormsModule],
    template: `
    <aside class="flex flex-col w-64 min-h-screen bg-surface-raised border-r border-slate-800 shrink-0">
      <!-- Logo -->
      <div class="flex items-center gap-3 px-5 py-4 border-b border-slate-800">
        <div class="w-8 h-8 rounded-lg bg-primary-600 flex items-center justify-center text-white font-bold text-sm">A</div>
        <span class="text-slate-100 font-semibold tracking-wide">APM</span>
      </div>

      <!-- Projects -->
      <div class="flex-1 overflow-y-auto py-3">
        <div class="flex items-center justify-between px-4 mb-2">
          <span class="text-xs font-semibold text-slate-500 uppercase tracking-widest">Projects</span>
          <button (click)="toggleNewForm()" class="text-slate-500 hover:text-primary-400 transition-colors text-lg leading-none" title="New project">+</button>
        </div>

        <!-- New project inline form -->
        @if (showNewForm()) {
          <div class="mx-3 mb-3 p-3 bg-surface-overlay rounded-xl border border-slate-700 animate-fade-in">
            <input
              id="new-project-name"
              class="input-base mb-2 text-xs"
              placeholder="Project name"
              [(ngModel)]="newName"
              (keydown.enter)="createProject()"
              (keydown.escape)="cancelNewForm()"
              autofocus
            />
            <div class="flex items-center gap-2 mb-2">
              <label class="text-xs text-slate-400">Color</label>
              <input type="color" [(ngModel)]="newColor" class="w-6 h-6 rounded cursor-pointer border-none bg-transparent" />
            </div>
            <div class="flex gap-2">
              <button (click)="createProject()" class="btn-primary text-xs py-1 px-3 flex-1">Create</button>
              <button (click)="cancelNewForm()" class="btn-ghost text-xs py-1 px-3">Cancel</button>
            </div>
          </div>
        }

        <!-- Project list -->
        @if (projects().length === 0 && !loading()) {
          <p class="px-4 text-xs text-slate-600 italic mt-2">No projects yet</p>
        }
        @for (project of projects(); track project.id) {
          <button
            class="w-full flex items-center gap-3 px-4 py-2.5 hover:bg-surface-overlay rounded-lg mx-1 transition-colors text-left group"
            [class.bg-surface-overlay]="activeProjectId() === project.id"
            (click)="selectProject(project)"
          >
            <span class="w-2.5 h-2.5 rounded-full shrink-0" [style.background-color]="project.color"></span>
            <span class="text-sm text-slate-300 group-hover:text-slate-100 truncate flex-1"
                  [class.text-slate-100]="activeProjectId() === project.id">
              {{ project.name }}
            </span>
          </button>
        }
      </div>

      <!-- Footer -->
      <div class="px-4 py-3 border-t border-slate-800">
        <p class="text-xs text-slate-600">Local · No cloud</p>
      </div>
    </aside>
  `,
})
export class SidebarComponent implements OnInit {
    private projectSvc = inject(ProjectService);

    projects = signal<Project[]>([]);
    loading = signal(true);
    activeProjectId = signal<number | null>(null);
    showNewForm = signal(false);
    newName = '';
    newColor = '#6366f1';

    ngOnInit(): void {
        this.loadProjects();
    }

    loadProjects(): void {
        this.projectSvc.list().subscribe({
            next: (data) => { this.projects.set(data); this.loading.set(false); },
            error: () => this.loading.set(false),
        });
    }

    toggleNewForm(): void { this.showNewForm.set(true); }
    cancelNewForm(): void { this.showNewForm.set(false); this.newName = ''; }

    createProject(): void {
        if (!this.newName.trim()) return;
        this.projectSvc.create({ name: this.newName.trim(), color: this.newColor }).subscribe({
            next: (p) => {
                this.projects.update(list => [...list, p]);
                this.cancelNewForm();
            },
        });
    }

    selectProject(project: Project): void {
        this.activeProjectId.set(project.id);
        // Emit a signal / event for the parent to consume — Phase 3 will use router params
    }
}
