import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { fetchProjectById } from "../api/project";
import { addTaskToProject, updateTaskStatus } from "../api/task";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import DraggableTask from "./DraggableTask";
import { useDrop } from "react-dnd";
import "../styles/TaskDashboard.css";
import { Task } from "../types/task";
import { normaliseStatus } from "../utils/normalise";

export const TaskDashboard: React.FC = () => {
  const { projectId } = useParams<{ projectId: string }>();
  const [projectName, setProjectName] = useState<string>("");
  const [tasks, setTasks] = useState<Task[]>([]);
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const [newTaskTitle, setNewTaskTitle] = useState<string>("");
  const [newTaskDescription, setNewTaskDescription] = useState<string>("");

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

  const handleAddTask = async () => {
    if (!newTaskTitle || !newTaskDescription) return;

    try {
      const rawTask = await addTaskToProject(newTaskTitle, newTaskDescription, 149, 63);

      const createdTask: Task = {
        ...rawTask,
        status: normaliseStatus(rawTask.status),
      };

      setTasks((prevTasks) => [...prevTasks, createdTask]);

      setIsModalOpen(false);
      setNewTaskTitle("");
      setNewTaskDescription("");
    } catch (error: any) {
      console.error("Failed to create task:", error?.response?.data || error.message);
    }
  };

  return (
    <div className="dashboard-container">
      <h1>{projectName}</h1>
      <div className="swimlanes">
        {(["to-do", "in-progress", "done"] as Task["status"][]).map((laneStatus) => {
          const [{ isOver, canDrop }, connectDropTarget] = useDrop(() => ({
            accept: "TASK",
            drop: async (item: { id: number }) => {
              //console.log("About to update task with id", item.id, "to status", laneStatus);

              try {
                const normalizedStatus = normaliseStatus(laneStatus);
                console.log(`Normalized status: ${normalizedStatus}`);
                const updatedTask = await updateTaskStatus(item.id, normalizedStatus);

                //console.log("Updated task:", updatedTask);

                setTasks((prevTasks) =>
                  prevTasks.map((task) =>
                    task.id === item.id ? { ...task, status: normalizedStatus } : task
                  )
                );
              } catch (error) {
                console.error("Failed to update task status:", error);
              }
            },
            canDrop: () => true,
            collect: (monitor) => ({
              isOver: monitor.isOver(),
              canDrop: monitor.canDrop(),
            }),
          }));

          return connectDropTarget(
            <div key={laneStatus} className={`swimlane ${isOver && canDrop ? "highlight" : ""}`}>
  <h2 className="swimlane-title">{laneStatus.replace("-", " ").toUpperCase()}</h2>
  
  {laneStatus === "to-do" && (
    <button className="add-task-btn" onClick={() => setIsModalOpen(true)}>
      <FontAwesomeIcon icon={faPlus} /> Add Task
    </button>
  )}

  {groupedTasks[laneStatus].map((task) => (
    <DraggableTask key={task.id} task={task} />
  ))}
</div>
          );
        })}
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
            <button className="add-task-btn" onClick={handleAddTask}>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="16"
    height="16"
    fill="none"
    stroke="currentColor"
    strokeWidth="2"
    strokeLinecap="round"
    strokeLinejoin="round"
    className="feather feather-plus"
    style={{ marginRight: 8 }}
  >
    <line x1="12" y1="5" x2="12" y2="19" />
    <line x1="5" y1="12" x2="19" y2="12" />
  </svg>
  Add Task
</button>
          </div>
        </div>
      )}
    </div>
  );
};