package handlers

import (
	"github.com/arkinjulijanto/go-base-api/boot/base"
	"github.com/arkinjulijanto/go-base-api/internal/repositories"
	"github.com/arkinjulijanto/go-base-api/internal/services"
)

type Handler struct {
	authService services.AuthService
}

type HandlerConfig struct {
	AuthService services.AuthService
}

func NewHandler(h *HandlerConfig) *Handler {
	return &Handler{
		authService: h.AuthService,
	}
}

func InitHandlers() *Handler {
	db := base.GetDBConn()

	userRepo := repositories.NewUserRepository((&repositories.URConfig{DB: db}))

	authService := services.NewAuthService(&services.ASConfig{UserRepo: userRepo})

	h := NewHandler(&HandlerConfig{
		AuthService: authService,
	})

	return h
}
