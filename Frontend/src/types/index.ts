export interface User {
  id: string;
  email: string;
  role: "employer" | "candidate" | "admin";
  is_banned: boolean;
  created_at: string;
}

export interface Company {
  id: string;
  owner_id: string;
  name: string;
  description: string;
  website: string;
  city: string;
  created_at: string;
}

export interface Job {
  id: string;
  company_id: string;
  title: string;
  description: string;
  salary_min: number;
  salary_max: number;
  type: "full" | "part" | "remote";
  status: "draft" | "open" | "closed";
  city: string;
  created_at: string;
}

export interface Application {
  id: string;
  job_id: string;
  candidate_id: string;
  resume_url: string;
  cover_letter: string;
  status: "pending" | "reviewed" | "interview" | "rejected" | "accepted";
  created_at: string;
}

export interface AuthResponse {
  access_token: string;
  user: User;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  role: "employer" | "candidate";
}
