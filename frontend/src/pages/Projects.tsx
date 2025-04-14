import React, { useState } from "react";
import { useAddProject, useFetchProject } from "../hooks/useProject";
import { useNavigate } from "react-router-dom";
import { User } from "../models/User";
import { Project } from "../models/Project"; 
import "../styles/Projects.scss";

const Projects: React.FC = () => {
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const [projectName, setProjectName] = useState<string>("");
  const [projectDescription, setProjectDescription] = useState<string>("");
  // TODO: fetch user
  const user: User = { id: 148, name: "Jane" }; 
  const { data: fetchedProjects, refetch } = useFetchProject(user.id);
  const { mutateAsync } = useAddProject(); 
  const navigate = useNavigate();



  const handleCreateProject = () => {
    setIsModalOpen(true);
  };

  const handleSubmitProject = async () => {
    if (projectName && projectDescription) {
      try {
        const newProject = {
          name: projectName,
          description: projectDescription,
          userId: user.id,
        };

        await mutateAsync(newProject); 
        setIsModalOpen(false); 
        setProjectName(""); 
        setProjectDescription("");
        await refetch();
      } catch (err) {
        console.error("Error adding project:", err);
      }
    }
  };

  return (
    <div className="projects-container">
      <h1>{user.name}'s Projects</h1>

      <div className="project-box">
        <button className="create-project-btn" onClick={handleCreateProject}>
          + Create Project
        </button>
      </div>

      <div className="projects-list">
      {fetchedProjects?.data?.sort((a: { created_at: string | number | Date; }, b: { created_at: string | number | Date; }) =>
  new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
).map((project: Project) => (
  <div
    key={project.id}
    className="project-card"
    onClick={() => navigate(`/dashboard/${project.id}`)}
  >
    <h2>{project.name}</h2>
    <p>{project.description}</p>
  </div>
))}
      </div>

      {isModalOpen && (
        <div className="modal">
          <div className="modal-content">
            <h2>Create New Project</h2>
            <input
              type="text"
              placeholder="Project name"
              value={projectName}
              onChange={(e) => setProjectName(e.target.value)}
            />
            <textarea
              placeholder="Project description"
              value={projectDescription}
              onChange={(e) => setProjectDescription(e.target.value)}
            />
            <button onClick={handleSubmitProject}>Submit</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default Projects;