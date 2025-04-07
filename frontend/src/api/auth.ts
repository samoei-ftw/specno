import axios from "axios";

const API_URL = import.meta.env.API_URL || "http://localhost:8080";

export const registerUser = async (email: string, password: string) => {
    const response = await axios.post(`${API_URL}/register`, { email, password });
    return response.data;
};

export const loginUser = async (email: string, password: string) => {
    const response = await axios.post(`${API_URL}/login`, { email, password });
    return response.data;
};