package response

import "github.com/arkinjulijanto/go-base-api/internal/models"

type RegisterResponse struct {
	Username string `json:"username"`
}

func (r *RegisterResponse) FormatResponse(u *models.User) {
	r.Username = u.Username
}
