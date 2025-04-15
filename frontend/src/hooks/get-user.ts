import { useQuery } from "@tanstack/react-query";
import { fetchUser as fetchUserAPI } from "../api/user";

export const useUserQuery = (userId: number) => {
  return useQuery({
    queryKey: ["user", userId], //userId is cacheable
    queryFn: () => fetchUserAPI(userId),
  });
};