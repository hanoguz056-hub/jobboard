package repository

import (
	"context"

	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(ctx context.Context, pool *pgxpool.Pool) (*UserRepository, error) {
	return &UserRepository{
		pool: pool,
	}, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	query := `INSERT INTO users (email, password_hash, role, is_banned) VALUES ($1, $2, $3, $4) RETURNING id, email, password_hash, role, is_banned, created_at`

	var u models.User

	err := r.pool.QueryRow(ctx, query, user.Email, user.PasswordHash, user.Role, user.IsBanned).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role, &u.IsBanned, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password_hash, role, is_banned, created_at FROM users WHERE email = $1`

	var u models.User 

	err := r.pool.QueryRow(ctx, query, email).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role, &u.IsBanned, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &u, err
}

func(r *UserRepository) GetAllUsers(ctx context.Context, limit, offset int) ([]models.User, error) {
	query := `SELECT id, email, password_hash, role, is_banned, created_at FROM users LIMIT $1 OFFSET $2`

	rows, err := r.pool.Query(ctx, query, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	
	users := make([]models.User, 0, limit)

	for rows.Next() {
		var u models.User

		err := rows.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role, &u.IsBanned, &u.CreatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, u)

		
	}

	if rows.Err() != nil {
			return nil, rows.Err()
		}

	return users, err
}

func(r *UserRepository) BanUser(ctx context.Context,userID string, isBanned bool)  error {
	query := `UPDATE users SET is_banned = $2 WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, userID, isBanned)

	if err != nil {
		return err
	}

	return nil 
}