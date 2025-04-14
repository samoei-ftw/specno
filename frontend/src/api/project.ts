import axios from "axios";
import { normaliseStatus } from "../utils/normalise";
const API_URL = "http://localhost:8082";

export const addProjectToUserAPI = async (
  name: string,
  description: string,
  user_id: number
) => {
  const response = await axios.post(`${API_URL}/projects`, {
    user_id,
    name,
    description,
  });
  return response.data;
};


export const fetchProjects = async (
    user_id: number
  ) => {
    const response = await axios.get(`${API_URL}/projects/${user_id}`);
    return response.data;
  };


  export const fetchProjectById = async (project_id: number) => {
    const response = await axios.get(`${API_URL}/projects?id=${project_id}`);
    const project = response.data;
  
    if (project.tasks) {
      project.tasks = project.tasks.map((task: any) => ({
        ...task,
        status: normaliseStatus(task.status),
      }));
    }
  
    return project;
  };