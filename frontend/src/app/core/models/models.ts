export interface Project {
    id: number;
    name: string;
    description?: string;
    color: string;
    icon?: string;
    createdAt?: string;
    updatedAt?: string;
    statuses?: Status[];
}

export interface Status {
    id: number;
    projectId: number;
    name: string;
    color: string;
    order: number;
}

export interface Tag {
    id: number;
    name: string;
    color: string;
}

export interface SubTask {
    id: number;
    taskId: number;
    title: string;
    isDone: boolean;
    order: number;
}

export type Priority = 'low' | 'medium' | 'high' | 'urgent';

export interface Task {
    id: number;
    projectId: number;
    statusId: number;
    title: string;
    description?: string;
    priority: Priority;
    dueDate?: string | null;
    startedAt?: string | null;
    completedAt?: string | null;
    createdAt?: string;
    updatedAt?: string;
    project?: Project;
    status?: Status;
    subTasks?: SubTask[];
    tags?: Tag[];
}

export interface TaskFilter {
    projectId?: number;
    statusId?: number;
    priority?: Priority;
    tagId?: number;
}
