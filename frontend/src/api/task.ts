import axios from "axios";

const API_URL = "http://localhost:8083";

type Task = {
    id: number;
    user_id: number;
    project_id: number;
    title: string;
    description: string;
    created_at: string;
    status: number;
  };
  
  type TaskResponse = {
    status: string;
    data: Task;
  };
  
  export const addTaskToProject = async (
    title: string,
    description: string,
    user_id: number,
    project_id: number,
  ): Promise<Task> => {
    const payload = {
      user_id,
      project_id,
      title,
      description,
    };
    const token = localStorage.getItem("token");
    if (!token) {
        throw new Error("No auth token found");
      }
  
    console.log("Sending task payload:", payload);
  
    try {
        const response = await axios.post<TaskResponse>(
          `${API_URL}/tasks`,
          payload,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );
        return response.data.data;
      } catch (error: any) {
        console.error("Error creating task:", error.response?.data || error.message);
        throw error;
      }
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