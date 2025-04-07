import { useMutation, UseMutationResult } from "@tanstack/react-query";
import axios, { AxiosResponse } from "axios";
import { addProjectToUserAPI } from "../api/project";

interface NewProject {
  name: string;
  description: string;
  userId: number;
}

  export const useAddProject = () => {
    return useMutation<AxiosResponse<NewProject>, Error, NewProject>({
        mutationFn: (newProject: NewProject) =>
        addProjectToUserAPI(newProject.name, newProject.description, newProject.userId),
    });
};