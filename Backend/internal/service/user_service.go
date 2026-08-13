package service

import (
	"context"
	"errors"

	"github.com/hanoguz056-hub/jobboard/internal/middleware"
	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/hanoguz056-hub/jobboard/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) (*UserService) {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, req models.RegisterRequest) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Email: req.Email,
		PasswordHash: string(hash),
		Role: req.Role,
	}

	return s.repo.CreateUser(ctx, user)
}



func(s *UserService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {

	foundUser, err := s.repo.GetUserByEmail(ctx, req.Email)
	
	if err != nil {
		return nil, err
	}

	if foundUser == nil || req.Password == "" {
		return nil, errors.New("Неверный email или пароль")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(req.Password))

	if err != nil {
		return nil, errors.New("Неверный email или пароль")
	}

	token, err := middleware.GenerateToken(foundUser.ID, string(foundUser.Role) )

	if err != nil {
		return nil, err 
	}

	return &models.AuthResponse{AccessToken: token, User: *foundUser}, nil
}

func(s *UserService) GetAllUsers(ctx context.Context, page, limit int) ([]models.User, error) {

	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	return s.repo.GetAllUsers(ctx, limit, offset)
}

func(s *UserService) BanUser(ctx context.Context, userID string, isBanned bool) error {
	
	return s.repo.BanUser(ctx, userID, isBanned)
}