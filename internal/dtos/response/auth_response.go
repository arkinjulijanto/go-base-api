package response

import "github.com/arkinjulijanto/go-base-api/internal/models"

type RegisterResponse struct {
	Username string `json:"username"`
}

func (r *RegisterResponse) FormatResponse(u *models.User) {
	r.Username = u.Username
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (l *LoginResponse) FormatResponse(u *models.User, token string) {
	l.Username = u.Username
	l.Token = token
}
