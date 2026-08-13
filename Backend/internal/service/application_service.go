package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/hanoguz056-hub/jobboard/internal/repository"
)

type ApplicationService struct {
	repo *repository.ApplicationRepository
}

func NewApplicationService(repo *repository.ApplicationRepository) (*ApplicationService) {
	return &ApplicationService{
		repo: repo,
	}
}

func(s *ApplicationService) Apply(ctx context.Context, jobID, candidateID, coverLetter string, fileBytes []byte, fileName string) (*models.Application, error) {
	
	ext := strings.ToLower(filepath.Ext(fileName))

	if ext != ".pdf" && ext != ".docx" {
		return nil, errors.New("Неверный формат файла")
	}

	uniqeName := uuid.New().String() + ext
	filePath := filepath.Join("./uploads/resumes", uniqeName)

	if err := os.WriteFile(filePath, fileBytes, 0644); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	resumeURL := "/uploads/resumes/" + uniqeName
	
	
	app := models.Application {
		JobID: jobID,
		CandidateID: candidateID,
		ResumeURL: resumeURL,
		CoverLetter: coverLetter,
		Status: "pending",
	}

	createdApp, err := s.repo.CreateApplication(ctx, app)

	if err != nil {
		_ = os.Remove(filePath)
		return nil, err
	}

	return createdApp, nil
}

func(s *ApplicationService) GetMyApplications(ctx context.Context, candidateID string) ([]models.Application, error) {
	jobs, err := s.repo.GetMyApplications(ctx, candidateID)

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func(s *ApplicationService) GetJobApplications(ctx context.Context, jobID string) ([]models.Application, error) {
	return s.repo.GetJobApplications(ctx, jobID)
}

func(s *ApplicationService) UpdateApplicationStatus(ctx context.Context, applicationID, status string) error {
	return s.repo.UpdateApplicationStatus(ctx, applicationID, status)
}