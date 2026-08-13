package repository

import (
	"context"
	"fmt"

	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type JobRepository struct {
	pool *pgxpool.Pool
}

func NewJobRepository(ctx context.Context, pool *pgxpool.Pool) (*JobRepository, error) {
	return &JobRepository{
		pool: pool,
	}, nil
}

func(r *JobRepository) GetJobs(ctx context.Context, limit, offset int, filters models.JobFilter) ([]models.Job, error) {
	query := `SELECT id, company_id, title, description, type, status, city, salary_min, salary_max, created_at FROM jobs WHERE 1=1`

	args := make([]interface{}, 0)

	argID := 1

	if filters.Search != "" {
		query += fmt.Sprintf(" AND title ILIKE $%d", argID)
		args = append(args, "%"+filters.Search+"%")
		argID++
	}

	if filters.City != "" {
		query += fmt.Sprintf(" AND city ILIKE $%d", argID)
		args = append(args, filters.City)
		argID++
	}

	if filters.Type != "" {
		query += fmt.Sprintf(" AND type = $%d", argID)
		args =append(args, filters.Type)
		argID++
	}

	if filters.Status != "" {
		query += fmt.Sprintf(" AND status = $%d", argID)
		args = append(args, filters.Status)
		argID++
	}

	if filters.SalaryMin > 0 {
		query += fmt.Sprintf(" AND salary_min >= $%d", argID)
		args = append(args, filters.SalaryMin)
		argID++
	}

	if filters.SalaryMax > 0 {
		query += fmt.Sprintf(" AND salary_max <= $%d", argID)
		args = append(args, filters.SalaryMax)
		argID++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argID, argID + 1)
	args = append(args, limit, offset)

	rows, err := r.pool.Query(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	jobs := make([]models.Job, 0, limit)

	
	for rows.Next() {
		var j models.Job

		err := rows.Scan(&j.ID, &j.CompanyID, &j.Title, &j.Description,&j.Type,&j.Status, &j.City, &j.SalaryMin, &j.SalaryMax, &j.CreatedAt)

		if err != nil {
			return nil, err
		}

		jobs = append(jobs, j)
	}

	

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
} 

func(r *JobRepository) GetJobByID(ctx context.Context, id string) (*models.Job, error) {
	var j models.Job

	query := `SELECT id, company_id, title, description, type, status, city, salary_min, salary_max, created_at FROM jobs WHERE id = $1`

	err := r.pool.QueryRow(ctx, query, id).Scan(&j.ID, &j.CompanyID, &j.Title, &j.Description, &j.Type, &j.Status, &j.City, &j.SalaryMin, &j.SalaryMax, &j.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &j, err
}

func(r *JobRepository) CreateJob(ctx context.Context,job models.Job) (*models.Job, error) {
	var j models.Job

	query := `INSERT INTO jobs (company_id, title, description, type, status, city, salary_min, salary_max) VALUES ($1, $2,$3, $4, $5, $6, $7, $8) RETURNING id, company_id, title, description, type, status, city, salary_min, salary_max, created_at`

	err := r.pool.QueryRow(ctx, query,job.CompanyID, job.Title, job.Description, job.Type, job.Status, job.City, job.SalaryMin, job.SalaryMax).Scan(&j.ID, &j.CompanyID,
	&j.Title, &j.Description, &j.Type, &j.Status, &j.City, &j.SalaryMin, &j.SalaryMax, &j.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &j, nil
}

func(r *JobRepository) UpdateJob(ctx context.Context, job models.Job) error {
	query := `UPDATE jobs SET title = $2, description = $3, type = $4, status = $5, city = $6, salary_min = $7, salary_max = $8 WHERE id = $1`

	_, err := r.pool.Exec(ctx, query,job.ID, job.Title, job.Description, job.Type, job.Status, job.City, job.SalaryMin, job.SalaryMax)

	if err != nil {
		return err
	}

	return nil
}

func(r *JobRepository) UpdateJobStatus(ctx context.Context, jobID, status string) error {
	query := `UPDATE jobs SET status = $2 WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, jobID, status)

	if err != nil {
		return err
	}

	return nil
}

func(r *JobRepository) DeleteJob(ctx context.Context, jobID string) error {
	query := `DELETE FROM jobs WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, jobID)

	if err != nil {
		return err
	}

	return nil
}

func(r *JobRepository) ApproveJob(ctx context.Context,jobID string) error {
	query := `UPDATE jobs SET status = "open" WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, jobID)

	if err != nil {
		return err
	}

	return nil
}