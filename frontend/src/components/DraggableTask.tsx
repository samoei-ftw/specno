import React, { useEffect, useState } from "react";
import { useDrag } from "react-dnd";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import "../styles/TaskDashboard.scss";

interface Task {
  id: number;
  title: string;
  description: string;
  status: "to-do" | "in-progress" | "done";
}

const DraggableTask = React.forwardRef<HTMLDivElement, { task: Task }>(({ task }, ref) => {
  const [{ isDragging }, drag] = useDrag(() => ({
    type: "TASK",
    item: { id: task.id, status: task.status },
    collect: (monitor) => ({
      isDragging: monitor.isDragging(),
    }),
  }));

  return (
    <div
      ref={(node) => {
        drag(node); 
        if (ref && typeof ref === "function") {
          ref(node)
        }
      }}
      className="task-card"
      style={{ opacity: isDragging ? 0.5 : 1 }}
    >
      <h3>{task.title}</h3>
      <p>{task.description}</p>
    </div>
  );
});

DraggableTask.displayName = "DraggableTask";

export default DraggableTask;