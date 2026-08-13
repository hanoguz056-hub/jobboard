package repository

import (
	"context"

	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepository struct {
	pool *pgxpool.Pool
}

func NewCompanyRepository(ctx context.Context, pool *pgxpool.Pool) (*CompanyRepository, error) {
	return &CompanyRepository{
		pool: pool,
	}, nil
}

func (r *CompanyRepository) GetCompanies(ctx context.Context, limit, offset int) ([]models.Company, error) {
	query := `SELECT id, owner_id, name, description, website, city, created_at FROM companies LIMIT $1 OFFSET $2`

	rows, err := r.pool.Query(ctx, query, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := make([]models.Company, 0, limit)

	for rows.Next() {
		var c models.Company

		err := rows.Scan(&c.ID, &c.OwnerID, &c.Name, &c.Description, &c.Website, &c.City, &c.CreatedAt)

		if err != nil {
			return nil, err
		}
		
		companies = append(companies, c)
	}

	if rows.Err() != nil {
			return nil, rows.Err()
		}

	return companies, nil
}


func(r *CompanyRepository) GetCompanyByID(ctx context.Context, id string) (*models.Company, error) {
	var c models.Company
	
	query := `SELECT id, owner_id, name, description, website, city, created_at FROM companies WHERE id = $1`

	err := r.pool.QueryRow(ctx, query, id).Scan(&c.ID, &c.OwnerID, &c.Name, &c.Description, &c.Website, &c.City, &c.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *CompanyRepository) CreateCompany(ctx context.Context, company models.Company) (*models.Company, error) {
	var c models.Company
	
	query := `INSERT INTO companies (owner_id, name, description, website, city) VALUES ($1, $2, $3, $4, $5) RETURNING id, owner_id, name, description, website, city, created_at`

	err := r.pool.QueryRow(ctx, query,company.OwnerID, company.Name, company.Description, company.Website, company.City).Scan(&c.ID, &c.OwnerID, &c.Name, &c.Description, &c.Website, &c.City, &c.CreatedAt)

	if err != nil {
		return nil, err
	}


	return &c, nil
}

func(r *CompanyRepository) UpdateCompany(ctx context.Context, company models.Company) error {
	query := `UPDATE companies SET name = $2, description = $3, website = $4, city = $5 WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, company.ID, company.Name, company.Description, company.Website, company.City)
	
	if err != nil {
		return err
	}

	return nil
}