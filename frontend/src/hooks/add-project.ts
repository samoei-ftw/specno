import { useMutation } from "@tanstack/react-query";
import { addProjectToUserAPI } from "../api/project";

interface NewProject {
  name: string;
  description: string;
  userId: number;
}

export const addProject = () => {
  return useMutation({
    mutationFn: (newProject: NewProject) =>
      addProjectToUserAPI(newProject.name, newProject.description, newProject.userId),
  });
};