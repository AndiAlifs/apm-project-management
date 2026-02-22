import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Status } from '../models/models';
import { environment } from '../../../environments/environment';

@Injectable({ providedIn: 'root' })
export class StatusService {
    private readonly base = `${environment.apiUrl}`;

    constructor(private http: HttpClient) { }

    listByProject(projectId: number): Observable<Status[]> {
        return this.http.get<Status[]>(`${this.base}/projects/${projectId}/statuses`);
    }

    create(projectId: number, payload: Partial<Status>): Observable<Status> {
        return this.http.post<Status>(`${this.base}/projects/${projectId}/statuses`, payload);
    }

    update(id: number, payload: Partial<Status>): Observable<Status> {
        return this.http.put<Status>(`${this.base}/statuses/${id}`, payload);
    }

    delete(id: number): Observable<void> {
        return this.http.delete<void>(`${this.base}/statuses/${id}`);
    }
}
