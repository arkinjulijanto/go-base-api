package routes

import (
	"github.com/arkinjulijanto/go-base-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, h *handlers.Handler) {
	r.POST("/register", h.Register)
}
