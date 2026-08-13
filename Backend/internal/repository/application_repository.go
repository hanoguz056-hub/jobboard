package repository

import (
	"context"

	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicationRepository struct {
	pool *pgxpool.Pool
}

func NewApplicationRepository(pool *pgxpool.Pool) (*ApplicationRepository) {
	return &ApplicationRepository{
		pool: pool,
	}
}


func(r *ApplicationRepository) CreateApplication(ctx context.Context, application models.Application) (*models.Application, error) {
	var a models.Application
	
	query := `INSERT INTO applications (job_id, candidate_id, resume_url, cover_letter, status) VALUES ($1, $2, $3, $4, $5) RETURNING id, job_id, candidate_id, resume_url, cover_letter, status, created_at`

	err := r.pool.QueryRow(ctx, query, application.JobID, application.CandidateID, application.ResumeURL, application.CoverLetter, application.Status).Scan(
		&a.ID, &a.JobID, &a.CandidateID, &a.ResumeURL, &a.CoverLetter, &a.Status, &a.CreatedAt)

	if err != nil {
		return nil, err
	}


	return &a, nil
}


func(r *ApplicationRepository) GetMyApplications(ctx context.Context, candidateID string) ([]models.Application, error) {
	query := `SELECT id, job_id, candidate_id, resume_url, cover_letter, status, created_at FROM applications WHERE candidate_id = $1`

	rows, err := r.pool.Query(ctx, query, candidateID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	myApplications := make([]models.Application, 0)

	for rows.Next() {
		var a models.Application

		err := rows.Scan(&a.ID, &a.JobID, &a.CandidateID, &a.ResumeURL, &a.CoverLetter, &a.Status, &a.CreatedAt)

		if err != nil {
			return nil, err
		}

		myApplications = append(myApplications, a)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return myApplications, nil
}

func(r *ApplicationRepository) GetJobApplications(ctx context.Context, jobID string) ([]models.Application, error) {
	query := `SELECT id, job_id, candidate_id, resume_url, cover_letter, status, created_at FROM applications WHERE job_id = $1`

	rows, err := r.pool.Query(ctx, query, jobID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	jobApplications := make([]models.Application, 0)

	for rows.Next() {
		var j models.Application

		err := rows.Scan(&j.ID, &j.JobID, &j.CandidateID, &j.ResumeURL, &j.CoverLetter, &j.Status, &j.CreatedAt)

		if err != nil {
			return nil, err
		}

		jobApplications = append(jobApplications, j)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return jobApplications, nil
}

func(r *ApplicationRepository) UpdateApplicationStatus(ctx context.Context,applicationID, status string) error {
	query := `UPDATE applications SET status = $2 WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, applicationID, status)

	if err != nil {
		return err
	}

	return nil
}

func(r *ApplicationRepository) GetApplicationByID(ctx context.Context, id string) (*models.Application, error) {
	var a models.Application
	query := `SELECT id, job_id, candidate_id, resume_url, cover_letter, status, created_at FROM applications WHERE id = $1`

	err := r.pool.QueryRow(ctx, query, id).Scan(&a.ID, &a.JobID, &a.CandidateID, &a.ResumeURL, &a.CoverLetter, &a.Status, &a.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

