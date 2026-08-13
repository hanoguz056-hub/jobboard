package service

import (
	"context"
	"errors"

	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/hanoguz056-hub/jobboard/internal/repository"
)

type CompanyService struct {
	repo *repository.CompanyRepository
}

func NewCompanyService(repo *repository.CompanyRepository) (*CompanyService) {
	return &CompanyService{
		repo: repo,
	}
}

func(s *CompanyService) GetCompanies(ctx context.Context, page, limit int) ([]models.Company, error) {
	offset := (page - 1) * limit

	return  s.repo.GetCompanies(ctx, limit, offset) 

}

func(s *CompanyService) GetCompanyByID(ctx context.Context, id string) (*models.Company, error) {
	company, err := s.repo.GetCompanyByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return company, nil
}

func(s *CompanyService) CreateCompany(ctx context.Context, company models.Company, ownerID string) (*models.Company, error) {
	company.OwnerID = ownerID 

	newCompany, err := s.repo.CreateCompany(ctx, company)

	if err != nil {
		return nil, err
	}

	return newCompany, nil
}

func(s *CompanyService) UpdateCompany(ctx context.Context, company models.Company, userID string) error {
	getCompany, err := s.repo.GetCompanyByID(ctx, company.ID)

	if err != nil {
		return errors.New("Компания не найдена")
	}

	if getCompany.OwnerID != userID {
		return errors.New("У вас нету доступа")
	}

	return s.repo.UpdateCompany(ctx, company)
}