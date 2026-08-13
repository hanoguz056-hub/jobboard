import api from "./axios";
import type { Job } from "../types";

export const getJobs = async (): Promise<Job[]> => {
  const response = await api.get("/jobs");
  return response.data;
};

export const getJobByID = async (id: string): Promise<Job> => {
  const response = await api.get(`/jobs/${id}`);
  return response.data;
};

export const createJob = async (data: Job): Promise<Job> => {
  const response = await api.post("/jobs", data);
  return response.data;
};

export const updateJobStatus = async (
  id: string,
  status: string,
): Promise<Job> => {
  const response = await api.patch(`/jobs/${id}/status`, { status });
  return response.data;
};

export const deleteJob = async (id: string): Promise<void> => {
  await api.delete(`/jobs/${id}`);
};
