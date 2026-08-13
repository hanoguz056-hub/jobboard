import api from "./axios";
import type { Application } from "../types";

export const applyToJob = async (
  jobID: string,
  formData: FormData,
): Promise<Application> => {
  const response = await api.post(`/jobs/${jobID}/apply`, formData, {
    headers: { "Content-Type": "multipart/form-data" },
  });
  return response.data;
};

export const getMyApplications = async (): Promise<Application[]> => {
  const response = await api.get(`/applications/my`);
  return response.data;
};

export const getJobApplications = async (
  id: string,
): Promise<Application[]> => {
  const response = await api.get(`/jobs/${id}/applications`);
  return response.data;
};

export const updateApplicationStatus = async (
  id: string,
  status: string,
): Promise<Application> => {
  const response = await api.patch(`/applications/${id}/status`, { status });
  return response.data;
};
