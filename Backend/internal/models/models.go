package models

import "time"


type Role string
type JobStatus string 
type JobType string
type ApplicationStatus string



const (
	RoleEmployer  Role = "employer"
	RoleCandidate Role = "candidate"
	RoleAdmin Role = "admin"
)

const (
	JobDraft JobStatus = "draft"
	JobOpen JobStatus = "open"
	JobClosed JobStatus = "closed"
)

const (
	JobFull JobType = "full"
	JobPart JobType = "part"
	JobRemote JobType = "remote"
)

const (
	ApplicationPending ApplicationStatus = "pending"
	ApplicationReviewed ApplicationStatus = "reviewed"
	ApplicationInterview ApplicationStatus = "interview"
	ApplicationRejected ApplicationStatus = "rejected"
	ApplicationAccepted ApplicationStatus = "accepted"
)

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Role         Role `json:"role"`
	IsBanned     bool   `json:"is_banned"`
	CreatedAt    time.Time `json:"created_at"`
}

type Company struct {
	ID string `json:"id"`
	OwnerID string `json:"owner_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Website string `json:"website"`
	City string `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

type Job struct {
	ID string `json:"id"`
	CompanyID string `json:"company_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	SalaryMin float64 `json:"salary_min"`
	SalaryMax float64 `json:"salary_max"`
	Type JobType `json:"type"`
	Status JobStatus `json:"status"`
	City string `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

type Application struct {
	ID string `json:"id"`
	JobID string `json:"job_id"`
	CandidateID string `json:"candidate_id"`
	ResumeURL string `json:"resume_url"`
	CoverLetter string `json:"cover_letter"`
	Status ApplicationStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
    User User `json:"user"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type RegisterRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Role Role `json:"role"`
}


type JobFilter struct {
	Search string
	City string
	Type string
	Status JobStatus
	SalaryMin float64
	SalaryMax float64
	Page int
	Limit int
}

type ApplicationRequest struct {
	JobID string `json:"job_id"`
	CoverLetter string `json:"cover_letter"`
}