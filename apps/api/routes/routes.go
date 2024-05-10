package routes

import (
	"net/http"

	"github.com/arkinjulijanto/go-base-api/internal/handlers"
	"github.com/arkinjulijanto/go-base-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, h *handlers.Handler) {
	r.GET("/ping", func(c *gin.Context) {
		logger.LogInfo("here is the ping request", c)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")

	api.POST("/register", h.Register)
}
