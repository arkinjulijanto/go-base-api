package services

import (
	"context"

	"github.com/arkinjulijanto/go-base-api/internal/models"
	"github.com/arkinjulijanto/go-base-api/internal/repositories"
	"github.com/arkinjulijanto/go-base-api/pkg/custom_error"
	"github.com/arkinjulijanto/go-base-api/pkg/jwt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(ctx context.Context, u *models.User) (*models.User, error)
	Login(ctx context.Context, username, password string) (*models.User, string, error)
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

func (a *authService) Login(ctx context.Context, username, password string) (*models.User, string, error) {
	user, err := a.userRepo.FindByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "", custom_error.NewErrorWrapper(custom_error.CodeClientUnauthorized, "invalid credentials", nil, err)
		}
		return nil, "", custom_error.NewErrorWrapper(custom_error.CodeServerError, "failed to get user data", nil, err)
	}

	if user.Password != password {
		return nil, "", custom_error.NewErrorWrapper(custom_error.CodeClientUnauthorized, "invalid credentials", nil, err)
	}

	jwtToken, err := generateJWT(*user)
	if err != nil {
		return nil, "", custom_error.NewErrorWrapper(custom_error.CodeServerError, "failed to generate jwt", nil, err)
	}

	return user, jwtToken, nil
}

func generateJWT(user models.User) (string, error) {
	jwtToken, err := jwt.GenerateJWTToken(user)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
