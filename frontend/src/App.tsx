import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Register from "./pages/Register";
import Projects from "./pages/Projects";
import TaskDashboard from "./components/TaskDashboard";
import { useLocation } from "react-router-dom";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Register />} />
                <Route path="/projects" element={<Projects />} />
                <Route path="/task-dashboard/id" element={<TaskDashboard projectName={""} initialTasks={[]} />} />
            </Routes>
        </Router>
    );
}
const TaskDashboardWrapper = () => {
    const location = useLocation();
    const { projectName, tasks } = location.state || {};
    return <TaskDashboard projectName={projectName} initialTasks={tasks} />;
  };
export default App;