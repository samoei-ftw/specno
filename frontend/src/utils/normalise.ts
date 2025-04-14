import { TaskStatus } from "../types/task";

// API → Frontend
export const normaliseStatus = (apiStatus: string): TaskStatus => {
  const normalized = apiStatus.toLowerCase().replace(" ", "-");
  if (["to-do", "in-progress", "done"].includes(normalized)) {
    return normalized as TaskStatus;
  }
  throw new Error(`Unknown status from API: ${apiStatus}`);
};

// Frontend → Backend
export const denormaliseStatus = (status: TaskStatus): number => {
  switch (status) {
    case "to-do":
      return 0;
    case "in-progress":
      return 1;
    case "done":
      return 2;
    default:
      throw new Error(`Unknown status for backend: ${status}`);
  }
};