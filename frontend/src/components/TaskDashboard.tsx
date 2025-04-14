import React, { useEffect, useState, useRef } from "react";
import { useParams } from "react-router-dom";
import { fetchProjectById } from "../api/project";
import { addTaskToProject } from "../api/task";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import DraggableTask from "./DraggableTask";
import { useDrop } from "react-dnd";
import "../styles/TaskDashboard.css";


interface Task {
  id: number;
  title: string;
  description: string;
  status: "to-do" | "in-progress" | "done";
}

export const TaskDashboard: React.FC = () => {
  const { projectId } = useParams<{ projectId: string }>();
  const [projectName, setProjectName] = useState<string>("");
  const [tasks, setTasks] = useState<Task[]>([]);
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const [newTaskTitle, setNewTaskTitle] = useState<string>("");
  const [newTaskDescription, setNewTaskDescription] = useState<string>("");
  const [taskStatus, setTaskStatus] = useState<"to-do" | "in-progress" | "done">("to-do");

  useEffect(() => {
    if (projectId) {
      fetchProjectById(Number(projectId))
        .then((project) => {
          setProjectName(project.name);
          setTasks(project.tasks || []);
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

  const [{ isOver, canDrop }, drop] = useDrop(() => ({
    accept: "TASK",
    drop: (item: { id: number }) => {
      const updatedTasks = tasks.map((task) => {
        if (task.id === item.id) {
          return { ...task, status: taskStatus };
        }
        return task;
      });

      setTasks(updatedTasks);
    },
    canDrop: (item: { id: number }) => {
      return true; 
    },
    collect: (monitor) => ({
      isOver: monitor.isOver(),
      canDrop: monitor.canDrop(),
    }),
  }));

  const handleAddTask = async () => {
    if (newTaskTitle && newTaskDescription) {
      try {
        const createdTask = await addTaskToProject(
          newTaskTitle,
          newTaskDescription,
          148,  //TODO
          62 
        );
  
        // append the newly created task to the local task list
        setTasks([...tasks, createdTask]);
  
        // reset form state
        setIsModalOpen(false);
        setNewTaskTitle("");
        setNewTaskDescription("");
      } catch (error) {
        console.error("Failed to create task:", error);
      }
    }
  };

  const dropRef = useRef<HTMLDivElement>(null);

  return (
    <div className="dashboard-container">
      <h1>{projectName}</h1>
      <div className="swimlanes">
  {["to-do", "in-progress", "done"].map((lane) => (
    <div
      key={lane}
      className="swimlane"
      onClick={() => setTaskStatus(lane as Task["status"])}
    >
      <div className="swimlane-header">
        <h2>{lane.replace("-", " ").toUpperCase()}</h2>
        <button
          className="add-task-btn"
          onClick={() => {
            setTaskStatus(lane as Task["status"]);
            setIsModalOpen(true);
          }}
        >
          <FontAwesomeIcon icon={faPlus} />
        </button>
      </div>
      <div className="tasks">
        {groupedTasks[lane as keyof typeof groupedTasks].map((task) => (
          <DraggableTask key={task.id} task={task} />
        ))}
      </div>
    </div>
  ))}
</div>

      {isModalOpen && (
        <div className="modal">
          <div className="modal-content">
            <h2>Add New Task</h2>
            <input
              type="text"
              placeholder="Task title"
              value={newTaskTitle}
              onChange={(e) => setNewTaskTitle(e.target.value)}
            />
            <textarea
              placeholder="Task description"
              value={newTaskDescription}
              onChange={(e) => setNewTaskDescription(e.target.value)}
            />
            <button onClick={handleAddTask}>Add Task</button>
          </div>
        </div>
      )}
    </div>
  );
};