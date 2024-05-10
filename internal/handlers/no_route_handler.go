package handlers

import (
	"github.com/arkinjulijanto/go-base-api/internal/utils/json_util"
	"github.com/arkinjulijanto/go-base-api/pkg/custom_error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) NoRoute(c *gin.Context) {
	httpres := json_util.ResponseError(c, custom_error.NewNotFoundError("path not found"))
	c.JSON(httpres.Code, httpres)
}
