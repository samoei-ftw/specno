import { fetchProjects as fetchProjectsAPI } from "../api/project";
import { useMutation } from "@tanstack/react-query";

export const useFetchUser = () => {
  return useMutation({
    mutationFn: ({ userId }: { userId: number }) => fetchProjectsAPI(userId),
  });
};