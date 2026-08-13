import api from "./axios";
import type { Company } from "../types";

export const getCompanies = async (): Promise<Company[]> => {
  const response = await api.get("/companies");
  return response.data;
};

export const getCompanyByID = async (id: string): Promise<Company> => {
  const response = await api.get(`/companies/${id}`);
  return response.data;
};

export const createCompany = async (data: Company): Promise<Company> => {
  const response = await api.post("/companies", data);
  return response.data;
};

export const updateCompany = async (
  id: string,
  data: Company,
): Promise<Company> => {
  const response = await api.put(`/companies/${id}`, data);
  return response.data;
};
