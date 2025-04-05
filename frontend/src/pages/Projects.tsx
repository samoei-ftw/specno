import React, { useState } from "react";
import { useAddProject } from "../hooks/usePage"; // make sure the hook path is correct
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

  // Use the hook to handle adding projects
  const { mutate, isError, error, data } = useAddProject(); // Correctly typed hook return

  const handleCreateProject = () => {
    setIsModalOpen(true); 
  };

  const handleSubmitProject = () => {
    if (projectName && projectDescription) {
      const newProject = { name: projectName, description: projectDescription };

      setProjects((prevProjects) => [...prevProjects, newProject]);

      const userId = 146; // Replace with dynamic user ID

      mutate({
        name: projectName,
        description: projectDescription,
        userId,
      });

      setIsModalOpen(false); 
      setProjectName(""); 
      setProjectDescription("");
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

      {/* Display projects */}
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

      {/* Error handling */}
      {isError && <p className="error-text">Error: {error?.message}</p>}

      {/* Success message */}
      {data && <p className="success-text">Project added successfully!</p>}
    </div>
  );
};

export default Projects;