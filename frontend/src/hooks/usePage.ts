import { useMutation, UseMutationResult } from "@tanstack/react-query";
import { addProjectToUserAPI } from "../api/project";

// Define the project type for better typing
interface AddProjectParams {
  name: string;
  description: string;
  userId: number;
}

export const useAddProject = (): UseMutationResult<any, Error, AddProjectParams> => {
  return useMutation({
    mutationFn: ({ name, description, userId }: AddProjectParams) =>
      addProjectToUserAPI(name, description, userId),
  });
};