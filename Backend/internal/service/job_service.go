package service

import (
	"context"
	"errors"

	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/hanoguz056-hub/jobboard/internal/repository"
)

type JobService struct {
	repo *repository.JobRepository
	companyRepo * repository.CompanyRepository
}

func NewJobService(repo *repository.JobRepository, companyRepo *repository.CompanyRepository) (*JobService) {
	return &JobService{
		repo: repo,
		companyRepo: companyRepo,
	}
}


func(s *JobService) GetJobs(ctx context.Context, filters models.JobFilter) ([]models.Job, error) {
	offset := (filters.Page - 1) * filters.Limit

	return s.repo.GetJobs(ctx, filters.Limit, offset, filters)
} 

func(s *JobService) GetJobByID(ctx context.Context, id string) (*models.Job, error) {
	return s.repo.GetJobByID(ctx, id)
}

func(s *JobService) CreateJob(ctx context.Context,job models.Job, companyID string) (*models.Job, error) {
	job.CompanyID = companyID
	job.Status = models.JobDraft

	return s.repo.CreateJob(ctx, job)
}

func(s *JobService) UpdateJob(ctx context.Context, job models.Job, userID string) error {
	existingJob, err := s.repo.GetJobByID(ctx, job.ID)

	if err != nil {
		return errors.New("Вакансия не найдена")
	}

	company, err := s.companyRepo.GetCompanyByID(ctx, existingJob.CompanyID)

	if err != nil {
		return errors.New("Компания не найдена")
	}

	if company.OwnerID != userID {
		return errors.New("Нет доступа")
	}	

	return s.repo.UpdateJob(ctx, job)
}

func(s *JobService) UpdateJobStatus(ctx context.Context, jobID, status, userID string) error {
	existingJob, err := s.repo.GetJobByID(ctx, jobID)

	if err != nil {
		return errors.New("Вакансия не найдена")
	}

	company, err := s.companyRepo.GetCompanyByID(ctx, existingJob.CompanyID)

	if err != nil {
		return errors.New("Компания не найдена")
	}

	if company.OwnerID != userID {
		return errors.New("Нет доступа к вакансии")
	}

	return s.repo.UpdateJobStatus(ctx, jobID, status)
}

func(s *JobService) DeleteJob(ctx context.Context, jobID, userID, userRole string) error {
	existingJob, err := s.repo.GetJobByID(ctx, jobID)

	if err != nil {
		return errors.New("Вакансия не найдена")
	}

	if userRole == "admin" {
		return s.repo.DeleteJob(ctx, jobID)
	}

	company, err := s.companyRepo.GetCompanyByID(ctx, existingJob.CompanyID)

	if err != nil {
		return errors.New("Компания не найдена")
	}

	if company.OwnerID != userID {
		return errors.New("У вас нету доступа")
	}

	return s.repo.DeleteJob(ctx,jobID)
}

func(s *JobService) ApproveJob(ctx context.Context, jobID string) error {
	return s.repo.ApproveJob(ctx, jobID)
}