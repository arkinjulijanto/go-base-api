package handlers

import (
	"context"
	"net/http"

	"github.com/arkinjulijanto/go-base-api/internal/dtos/request"
	"github.com/arkinjulijanto/go-base-api/internal/dtos/response"
	"github.com/arkinjulijanto/go-base-api/internal/utils/json_util"
	"github.com/arkinjulijanto/go-base-api/pkg/custom_error"
	"github.com/arkinjulijanto/go-base-api/pkg/validator"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", c.GetHeader("Request-Id"))

	var req request.RegisterRequest
	var res response.RegisterResponse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		validation := validator.FormatValidation(err)
		httpres := json_util.ResponseError(c, custom_error.NewUnprocessibleEntityError(validation))
		c.JSON(httpres.Code, httpres)
		return
	}

	u := req.ConvertToModel()
	user, err := h.authService.Register(ctx, u)
	if err != nil {
		httpres := json_util.ResponseError(c, err)
		c.JSON(httpres.Code, httpres)
		return
	}

	res.FormatResponse(user)

	httpres := json_util.ResponseSuccess(c, res, "register success", http.StatusCreated)
	c.JSON(httpres.Code, httpres)
}

func (h *Handler) Login(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestId", c.GetHeader("Request-Id"))

	var req request.LoginRequest
	var res response.LoginResponse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		validation := validator.FormatValidation(err)
		httpres := json_util.ResponseError(c, custom_error.NewUnprocessibleEntityError(validation))
		c.JSON(httpres.Code, httpres)
		return
	}

	u := req.ConvertToModel()
	user, token, err := h.authService.Login(ctx, u.Username, u.Password)
	if err != nil {
		httpres := json_util.ResponseError(c, err)
		c.JSON(httpres.Code, httpres)
		return
	}

	res.FormatResponse(user, token)
	httpres := json_util.ResponseSuccess(c, res, "login success", http.StatusOK)
	c.JSON(httpres.Code, httpres)
}
