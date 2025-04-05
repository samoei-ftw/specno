import React, { useState } from "react";
import { useRegister } from "../hooks/useRegister";
import { useNavigate } from "react-router-dom";
import "../styles/Register.css";

const FadeButton: React.FC<{ title: string; onClick: () => void }> = ({ title, onClick }) => {
    const [hover, setHover] = useState(false);

    return (
        <button
            className={`fade-button ${hover ? "hover" : ""}`}
            onMouseEnter={() => setHover(true)}
            onMouseLeave={() => setHover(false)}
            onClick={onClick}
        >
            {title}
        </button>
    );
};

const Register = () => {
    const navigate = useNavigate();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const { mutate, isPending, error, data } = useRegister();

    const handleSubmit = () => {
        mutate({ email, password });
    };

    const handleLogin = () => {
        navigate("/projects");
    };

    return (
        <div className="register-container">
            <h1 className="register-title">Tasko</h1>
            <input
                type="email"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="register-input"
            />
            <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="register-input"
            />
            <FadeButton title={isPending ? "Registering..." : "Register"} onClick={handleSubmit} />
            <FadeButton title="Login" onClick={handleLogin} />
            {error && <p className="error-text">Error: {error.message}</p>}
            {data && <p className="success-text">Registered successfully!</p>}
        </div>
    );
};

export default Register;