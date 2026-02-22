import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Tag } from '../models/models';
import { environment } from '../../../environments/environment';

@Injectable({ providedIn: 'root' })
export class TagService {
    private readonly base = `${environment.apiUrl}/tags`;
    private readonly tasksBase = `${environment.apiUrl}/tasks`;

    constructor(private http: HttpClient) { }

    list(): Observable<Tag[]> {
        return this.http.get<Tag[]>(this.base);
    }

    create(payload: Partial<Tag>): Observable<Tag> {
        return this.http.post<Tag>(this.base, payload);
    }

    update(id: number, payload: Partial<Tag>): Observable<Tag> {
        return this.http.put<Tag>(`${this.base}/${id}`, payload);
    }

    delete(id: number): Observable<void> {
        return this.http.delete<void>(`${this.base}/${id}`);
    }

    assign(taskId: number, tagId: number): Observable<void> {
        return this.http.post<void>(`${this.tasksBase}/${taskId}/tags/${tagId}`, {});
    }

    remove(taskId: number, tagId: number): Observable<void> {
        return this.http.delete<void>(`${this.tasksBase}/${taskId}/tags/${tagId}`);
    }
}
