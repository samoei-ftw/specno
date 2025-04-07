import React, { useState } from "react";
import { useAddProject } from "../hooks/usePage";
import "../styles/Projects.css";

interface Project {
  name: string;
  description: string;
}

const Projects: React.FC = () => {
  const [projects, setProjects] = useState<Project[]>([]); // project states
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false); // toggle modal visibility
  const [projectName, setProjectName] = useState<string>("");
  const [projectDescription, setProjectDescription] = useState<string>("");

  const { mutateAsync, isError, error, data } = useAddProject();

  const handleCreateProject = () => {
    setIsModalOpen(true);
  };

  const handleSubmitProject = async () => {
    if (projectName && projectDescription) {
      const userId = 146; // You can dynamically get this value as needed

      try {
        const newProject = {
          name: projectName,
          description: projectDescription,
          userId,
        };

        const token = localStorage.getItem("token");
        if (!token) {
          console.error("No token found, cannot add project.");
          return;
        }

        // Call mutateAsync with new project
        const response = await mutateAsync(newProject);

        console.log("Project added successfully:", response);
        setProjects((prev) => [...prev, { name: projectName, description: projectDescription }]);
        setIsModalOpen(false);
        setProjectName("");
        setProjectDescription("");
      } catch (err) {
        // Log the error for debugging
        console.error("Error adding project:", err);
      }
    }
  };

  return (
    <div className="projects-container">
      <h1>My Projects</h1>

      <div className="project-box">
        <button className="create-project-btn" onClick={handleCreateProject}>
          + Create Project
        </button>
      </div>

      <div className="projects-list">
        {projects.map((project, index) => (
          <div key={index} className="project-block">
            <h3>{project.name}</h3>
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
              placeholder="Project Name"
              value={projectName}
              onChange={(e) => setProjectName(e.target.value)}
              className="project-input"
            />
            <textarea
              placeholder="Project Description"
              value={projectDescription}
              onChange={(e) => setProjectDescription(e.target.value)}
              className="project-input"
            />
            <button className="done-btn" onClick={handleSubmitProject}>
              Done
            </button>
            <button
              className="close-modal-btn"
              onClick={() => setIsModalOpen(false)}
            >
              Cancel
            </button>
          </div>
        </div>
      )}

      {isError && <p className="error-text">Error: {error?.message}</p>}

      {data && <p className="success-text">Project added successfully!</p>}
    </div>
  );
};

export default Projects;