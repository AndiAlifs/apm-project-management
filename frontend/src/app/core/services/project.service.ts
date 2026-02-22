import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Project } from '../models/models';
import { environment } from '../../../environments/environment';

@Injectable({ providedIn: 'root' })
export class ProjectService {
    private readonly base = `${environment.apiUrl}/projects`;

    constructor(private http: HttpClient) { }

    list(): Observable<Project[]> {
        return this.http.get<Project[]>(this.base);
    }

    create(payload: Partial<Project>): Observable<Project> {
        return this.http.post<Project>(this.base, payload);
    }

    update(id: number, payload: Partial<Project>): Observable<Project> {
        return this.http.put<Project>(`${this.base}/${id}`, payload);
    }

    delete(id: number): Observable<void> {
        return this.http.delete<void>(`${this.base}/${id}`);
    }
}
