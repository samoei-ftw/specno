import { useState } from "react";
import { useRegister } from "../hooks/useRegister";
import React from "react";

const Register = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const { mutate, isPending, error, data } = useRegister();

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        mutate({ email, password });
    };

    return (
        <div className="flex h-screen items-center justify-center bg-black">
            <form onSubmit={handleSubmit} className="p-6 bg-gray-800 rounded-lg shadow-md">
                <h2 className="text-white text-xl mb-4">Register</h2>
                <input
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    className="mb-2 p-2 w-full rounded"
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="mb-2 p-2 w-full rounded"
                />
                <button type="submit" className="w-full p-2 bg-blue-600 text-white rounded" disabled={isPending}>
                    {isPending ? "Registering..." : "Register"}
                </button>
                {error && <p className="text-red-500 mt-2">Error: {error.message}</p>}
                {data && <p className="text-green-500 mt-2">Registered successfully!</p>}
            </form>
        </div>
    );
};

export default Register;