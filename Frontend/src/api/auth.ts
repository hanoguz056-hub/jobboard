import api from "./axios";
import type { AuthResponse, LoginRequest, RegisterRequest } from "../types";

export const login = async (data: LoginRequest): Promise<AuthResponse> => {
  const response = await api.post("/auth/login", data);
  return response.data;
};

export const register = async (
  data: RegisterRequest,
): Promise<AuthResponse> => {
  const response = await api.post("/auth/register", data);
  return response.data;
};
