import React from "react";
import "../styles/TaskDashboard.css";

interface TaskDashboardProps {
  projectName: string;
  initialTasks: {
    id: number;
    title: string;
    status: "to-do" | "in-progress" | "done";
  }[];
}

const TaskDashboard: React.FC<TaskDashboardProps> = ({ projectName, initialTasks }) => {
  const groupedTasks = {
    "to-do": initialTasks.filter((task) => task.status === "to-do"),
    "in-progress": initialTasks.filter((task) => task.status === "in-progress"),
    "done": initialTasks.filter((task) => task.status === "done"),
  };

  return (
    <div className="dashboard-container">
      <h1>{projectName}</h1>
      <div className="swimlanes">
        {["to-do", "in-progress", "done"].map((lane) => (
          <div key={lane} className="swimlane">
            <h2>{lane.replace("-", " ").toUpperCase()}</h2>
            {groupedTasks[lane as keyof typeof groupedTasks].map((task) => (
              <div key={task.id} className="task-card">
                {task.title}
              </div>
            ))}
          </div>
        ))}
      </div>
    </div>
  );
};

export default TaskDashboard;