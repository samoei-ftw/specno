import React, { useState } from "react";
import { useRegister } from "../hooks/useRegister";
import { useLogin } from "../hooks/useLogin";
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

const ErrorModal: React.FC<{ errorMessage: string; onClose: () => void }> = ({
    errorMessage,
    onClose,
}) => (
    <div className="error-modal">
        <div className="modal-content">
            <h2>Error</h2>
            <p>{errorMessage}</p>
            <button onClick={onClose}>Close</button>
        </div>
    </div>
);

const Register = () => {
    const navigate = useNavigate();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [showErrorModal, setShowErrorModal] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const { mutate, isPending, error, data } = useRegister();
    const {
        mutate: mutateLogin,
        isPending: isLoginPending,
        error: loginError,
        data: loginData,
    } = useLogin();

    const handleSubmit = () => {
        mutate({ email, password });
    };

    const handleLogin = () => {
        mutateLogin(
            { email, password },
            {
                onSuccess: (response) => {
                    
                    //console.log('Login response:', response);
                    const token = response?.token;
                    //console.log('Token:', token);
                    
                    if (token) {
                        localStorage.setItem("token", token);
                        navigate("/projects");
                    } else {
                        console.error('Token not found in response');
                        setShowErrorModal(true);
                        setErrorMessage("Failed to retrieve token from response.");
                    }
                },
                onError: (error) => {
                    //console.error('Login error:', error);
                    setShowErrorModal(true);
                    setErrorMessage(error.message || "An error occurred while logging in.");
                },
            }
        );
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
            <FadeButton title={isLoginPending ? "Logging in..." : "Login"} onClick={handleLogin} />
            {(error || loginError) && (
                <p className="error-text">
                    Error: {(error || loginError)?.message}
                </p>
            )}
            {(data || loginData) && (
                <p className="success-text">
                    {data ? "Registered successfully!" : "Logged in!"}
                </p>
            )}
            
            {showErrorModal && (
                <ErrorModal
                    errorMessage={errorMessage}
                    onClose={() => setShowErrorModal(false)}
                />
            )}
        </div>
    );
};

export default Register;