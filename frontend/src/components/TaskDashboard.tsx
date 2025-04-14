import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { fetchProjectById } from "../api/project";
import { addTaskToProject, updateTaskStatus } from "../api/task";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import DraggableTask from "./DraggableTask";
import { useDrop } from "react-dnd";
import "../styles/TaskDashboard.css";
import { Task, TaskStatus } from "../types/task";
import { normaliseStatus } from "../utils/normalise";

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
              try {
                const updatedTask = await updateTaskStatus(item.id, laneStatus);
                setTasks((prevTasks) =>
                  prevTasks.map((task) =>
                    task.id === item.id ? { ...task, status: laneStatus } : task
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
          
          return (
            <div
              key={laneStatus}
              ref={connectDropTarget as unknown as React.RefObject<HTMLDivElement>} 
              className={`swimlane ${isOver && canDrop ? "highlight" : ""}`}
            >
              <div
  key={laneStatus}
  ref={connectDropTarget as unknown as React.RefObject<HTMLDivElement>}
  className={`swimlane ${isOver && canDrop ? "highlight" : ""}`}
>
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
            <button onClick={handleAddTask}>Add Task</button>
          </div>
        </div>
      )}
    </div>
  );
};