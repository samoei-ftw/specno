import { useQuery, useMutation, UseMutationResult } from "@tanstack/react-query";
import axios, { AxiosResponse } from "axios";
import { addProjectToUserAPI } from "../api/project";
import { fetchProjects as fetchProjectsAPI } from "../api/project";

export const getProject = (userId: number) => {
  return useQuery({
    queryKey: ["projects", userId],
    queryFn: () => fetchProjectsAPI(userId),
  });
};