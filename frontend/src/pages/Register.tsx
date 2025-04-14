import React, { useState } from "react";
import { useRegister } from "../hooks/useRegister";
import { useLogin } from "../hooks/useLogin";
import { useNavigate } from "react-router-dom";
import "../styles/AuthForm.scss";

const ErrorModal = ({ errorMessage, onClose }: { errorMessage: string; onClose: () => void }) => (
    <div className="error-modal">
        <div className="modal-content">
            <h2>Error</h2>
            <p>{errorMessage}</p>
            <button onClick={onClose}>Close</button>
        </div>
    </div>
);

export default function Register() {
    const navigate = useNavigate();
    const [credentials, setCredentials] = useState({ email: "", password: "" });
    const [showErrorModal, setShowErrorModal] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const { mutate, isPending, error, data } = useRegister();
    const {
        mutate: mutateLogin,
        isPending: isLoginPending,
        error: loginError,
        data: loginData,
    } = useLogin();

    const handleChange = (ev: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = ev.target;
        setCredentials({ ...credentials, [name]: value });
    };

    const handleRegister = () => {
        mutate(credentials);
    };

    const handleLogin = () => {
        mutateLogin(credentials, {
            onSuccess: (response) => {
                const token = response?.token;
                if (token) {
                    localStorage.setItem("token", token);
                    navigate("/projects");
                } else {
                    setShowErrorModal(true);
                    setErrorMessage("Failed to retrieve token from response.");
                }
            },
            onError: (error: any) => {
                setShowErrorModal(true);
                setErrorMessage(error.message || "An error occurred while logging in.");
            },
        });
    };

    return (
        <section className="auth-page">
            <form className="auth-form" onSubmit={(e) => e.preventDefault()}>
                <h1 className="auth-title">Tasko</h1>
                <input
                    type="email"
                    name="email"
                    placeholder="Email"
                    value={credentials.email}
                    onChange={handleChange}
                    required
                />
                <input
                    type="password"
                    name="password"
                    placeholder="Password"
                    value={credentials.password}
                    onChange={handleChange}
                    required
                />
                <button className="btn" onClick={handleRegister}>
                    {isPending ? "Registering..." : "Register"}
                </button>
                <button className="btn" onClick={handleLogin}>
                    {isLoginPending ? "Logging in..." : "Login"}
                </button>

                {(error || loginError) && (
                    <p className="error-text">
                        {(error || loginError)?.message}
                    </p>
                )}
                {(data || loginData) && (
                    <p className="success-text">
                        {data ? "Registered successfully!" : "Logged in!"}
                    </p>
                )}
            </form>

            {showErrorModal && (
                <ErrorModal errorMessage={errorMessage} onClose={() => setShowErrorModal(false)} />
            )}
        </section>
    );
}