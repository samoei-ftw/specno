import axios from "axios";
import { normaliseStatus } from "../utils/normalise";
import {Task, TaskResponse} from "../types/task";
const API_URL = "http://localhost:8083";

export const addTaskToProject = async (
  title: string,
  description: string,
  user_id: number,
  project_id: number,
): Promise<Task> => {
  const payload = { user_id, project_id, title, description };
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No auth token found");

  const response = await axios.post<TaskResponse>(
    `${API_URL}/tasks`,
    payload,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  const taskData = response.data.data;
  return {
    ...taskData,
    status: normaliseStatus(taskData.status),
  };
};


export const fetchTasks = async (
    project_id: number
  ) => {
    const response = await axios.get(`${API_URL}/tasks/${project_id}`);
    return response.data;
  };

export const fetchTaskById = async (
  task_id: number
) => {
  const response = await axios.get(`${API_URL}/tasks?id=${task_id}`);
  return response.data;
}

export const updateTaskStatus = async (taskId: number, statusInput: string) => {
  let status: number;

  switch (statusInput) {
    case "to-do":
      status = 0;
      break;
    case "in-progress":
      status = 1;
      break;
    case "done":
      status = 2;
      break;
    default:
      throw new Error("Invalid status value");
  }
    const response = await fetch(`http://localhost:8083/tasks/${taskId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: JSON.stringify({ status }),
    });
    
  
    if (!response.ok) {
      throw new Error("Failed to update task status");
    }
    else{
      console.log(`Successfully updated task status to ${status}`)
    }
  
    return response.json();
  };