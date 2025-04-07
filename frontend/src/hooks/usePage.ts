import { useMutation, UseMutationResult } from "@tanstack/react-query";
import axios, { AxiosResponse } from "axios";

interface NewProject {
  name: string;
  description: string;
  userId: number;
}

// Define the mutation function separately
const addProject = (newProject: NewProject): Promise<AxiosResponse<NewProject>> => {
  const token = localStorage.getItem("token");

  return axios.post("/api/projects", newProject, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};

export const useAddProject = () =>
  useMutation<AxiosResponse<NewProject>, Error, NewProject>({
    mutationFn: (newProject: NewProject) => {
      const token = localStorage.getItem("token");

      return axios.post("/api/projects", newProject, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
    },
  });