import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Task, TaskFilter } from '../models/models';
import { environment } from '../../../environments/environment';

@Injectable({ providedIn: 'root' })
export class TaskService {
    private readonly base = `${environment.apiUrl}`;

    constructor(private http: HttpClient) { }

    /** Global backlog with optional filters */
    list(filter?: TaskFilter): Observable<Task[]> {
        let params = new HttpParams();
        if (filter?.projectId) params = params.set('projectId', filter.projectId);
        if (filter?.statusId) params = params.set('statusId', filter.statusId);
        if (filter?.priority) params = params.set('priority', filter.priority);
        if (filter?.tagId) params = params.set('tagId', filter.tagId);
        return this.http.get<Task[]>(`${this.base}/tasks`, { params });
    }

    /** Tasks within a single project */
    listByProject(projectId: number): Observable<Task[]> {
        return this.http.get<Task[]>(`${this.base}/projects/${projectId}/tasks`);
    }

    getById(id: number): Observable<Task> {
        return this.http.get<Task>(`${this.base}/tasks/${id}`);
    }

    create(payload: Partial<Task>): Observable<Task> {
        return this.http.post<Task>(`${this.base}/tasks`, payload);
    }

    update(id: number, payload: Partial<Task>): Observable<Task> {
        return this.http.put<Task>(`${this.base}/tasks/${id}`, payload);
    }

    patchStatus(taskId: number, statusId: number): Observable<Task> {
        return this.http.patch<Task>(`${this.base}/tasks/${taskId}/status`, { statusId });
    }

    delete(id: number): Observable<void> {
        return this.http.delete<void>(`${this.base}/tasks/${id}`);
    }

    addSubTask(taskId: number, title: string): Observable<any> {
        return this.http.post(`${this.base}/tasks/${taskId}/subtasks`, { title });
    }

    updateSubTask(subtaskId: number, payload: { title?: string; isDone?: boolean }): Observable<any> {
        return this.http.put(`${this.base}/subtasks/${subtaskId}`, payload);
    }

    deleteSubTask(subtaskId: number): Observable<void> {
        return this.http.delete<void>(`${this.base}/subtasks/${subtaskId}`);
    }
}
