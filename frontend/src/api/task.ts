import axios from "axios";

const API_URL = "http://localhost:8083";

export const addTaskToProject = async (
  title: string,
  description: string,
  user_id: number,
  project_id: number,
) => {
    // TODO: response
  const response = await axios.post(`${API_URL}/tasks`, {
    user_id,
    name,
    description,
  });
  return response.data;
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