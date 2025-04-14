import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Register from "./pages/Register";
import Projects from "./pages/Projects";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import {TaskDashboard} from "./components/TaskDashboard";
import { useLocation } from "react-router-dom";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Register />} />
                <Route path="/projects" element={<Projects />} />
                <Route path="/dashboard/:projectId" element={<DndProvider backend={HTML5Backend}>
      <TaskDashboard />
    </DndProvider>} />
            </Routes>
        </Router>
    );
}
const TaskDashboardWrapper = () => {
    const location = useLocation();
    const { projectName, tasks } = location.state || {};
    return <TaskDashboard />;
  };
export default App;