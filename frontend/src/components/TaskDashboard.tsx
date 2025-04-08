import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import "../styles/TaskDashboard.css";
import { fetchProjectById } from "../api/project";

interface Task {
  id: number;
  title: string;
  status: "to-do" | "in-progress" | "done";
}

export const TaskDashboard: React.FC = () => {
  const { projectId } = useParams<{ projectId: string }>();
  const [projectName, setProjectName] = useState<string>("");
  const [tasks, setTasks] = useState<Task[]>([]);

  useEffect(() => {
    if (projectId) {
      fetchProjectById(Number(projectId))
        .then((project) => {
          setProjectName(project.name);
          setTasks(project.tasks || []);  // Ensure tasks is always an array
        })
        .catch((err) => {
          console.error("Error fetching project:", err);
        });
    }
  }, [projectId]);

  const groupedTasks = {
    "to-do": tasks.filter((task) => task.status === "to-do"),
    "in-progress": tasks.filter((task) => task.status === "in-progress"),
    "done": tasks.filter((task) => task.status === "done"),
  };

  return (
    <div className="dashboard-container">
      <h1>{projectName}</h1>
      <div className="swimlanes">
        {["to-do", "in-progress", "done"].map((lane) => (
          <div key={lane} className="swimlane">
            <h2>{lane.replace("-", " ").toUpperCase()}</h2>
            <div className="tasks">
              {groupedTasks[lane as keyof typeof groupedTasks].length === 0 ? (
                <p>No tasks</p>
              ) : (
                groupedTasks[lane as keyof typeof groupedTasks].map((task) => (
                  <div key={task.id} className="task-card">
                    {task.title}
                  </div>
                ))
              )}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};