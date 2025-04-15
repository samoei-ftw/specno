import { useMutation } from "@tanstack/react-query";
import { registerUser } from "../api/user";

export const useRegister = () => {
    return useMutation({
        mutationFn: ({ email, password }: { email: string; password: string }) =>
            registerUser(email, password),
    });
};