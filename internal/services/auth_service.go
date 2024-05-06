package services

import (
	"context"

	"github.com/arkinjulijanto/go-base-api/internal/models"
	"github.com/arkinjulijanto/go-base-api/internal/repositories"
	"github.com/arkinjulijanto/go-base-api/pkg/custom_error"
)

type AuthService interface {
	Register(ctx context.Context, u *models.User) (*models.User, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

type ASConfig struct {
	UserRepo repositories.UserRepository
}

func NewAuthService(a *ASConfig) AuthService {
	return &authService{
		userRepo: a.UserRepo,
	}
}

func (a *authService) Register(ctx context.Context, u *models.User) (*models.User, error) {
	user, err := a.userRepo.Create(u)
	if err != nil {
		return nil, custom_error.NewErrorWrapper(custom_error.CodeServerError, "failed to register", nil, err)
	}

	return user, nil
}
