package request

import "github.com/arkinjulijanto/go-base-api/internal/models"

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=2,max=20"`
	Password string `json:"password" binding:"required,min=5"`
}

func (r *RegisterRequest) ConvertToModel() *models.User {
	return &models.User{
		Username: r.Username,
		Password: r.Password,
	}
}
