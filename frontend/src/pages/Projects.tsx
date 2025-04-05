import React from "react";
import "../styles/Projects.css";

const Projects: React.FC = () => {
  return (
    <div className="projects-container">
      <div className="project-box">
        <h2 className="project-title">
          Create new project
        </h2>
        <button className="create-project-btn">
          + Create Project
        </button>
      </div>
    </div>
  );
};

export default Projects;