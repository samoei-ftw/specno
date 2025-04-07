import { fetchUser as fetchUserAPI } from "../api/user";
import { useMutation } from "@tanstack/react-query";

export const useFetchUser = () => {
  return useMutation({
    mutationFn: ({ id }: { id: number }) => fetchUserAPI(id),
  });
};