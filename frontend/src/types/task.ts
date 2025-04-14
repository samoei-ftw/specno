export type TaskStatus = "to-do" | "in-progress" | "done";

export interface Task {
  id: number;
  title: string;
  description: string;
  status: TaskStatus;
  user_id: number;
  project_id: number;
  created_at: string;
}

export interface TaskResponse {
  status: string;
  message: string;
  data: {
    id: number;
    title: string;
    description: string;
    status: string; // string bc from backend
    user_id: number;
    project_id: number;
    created_at: string;
  };
}