import { useMutation } from "@tanstack/react-query";
import { loginUser } from "../api/user";

export const login = () => {
  return useMutation({
    mutationFn: ({ email, password }: { email: string; password: string }) =>
      loginUser(email, password),
  });
};