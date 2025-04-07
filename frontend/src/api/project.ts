import axios from "axios";

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
    const response = await axios.post(`${API_URL}/projects?user_id=${user_id}`);
    return response.data;
  };