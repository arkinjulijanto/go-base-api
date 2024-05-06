package json_util

import (
	"errors"
	"net/http"

	"github.com/arkinjulijanto/go-base-api/internal/dtos/response"
	"github.com/arkinjulijanto/go-base-api/pkg/custom_error"
	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data interface{}, msg string, code ...int) response.JsonResponse {
	resCode := http.StatusOK
	if len(code) > 0 {
		resCode = code[0]
	}

	//TODO: add logger

	return responseJSON(data, msg, nil, resCode)
}

func ResponseError(c *gin.Context, err error) response.JsonResponse {
	resCode := http.StatusInternalServerError
	var ew *custom_error.ErrorWrapper
	if errors.As(err, &ew) {
		switch ew.Code {
		case custom_error.CodeClientError:
			resCode = http.StatusBadRequest
		case custom_error.CodeNotFoundError:
			resCode = http.StatusNotFound
		case custom_error.CodeConflictError:
			resCode = http.StatusConflict
		case custom_error.CodeClientUnauthorized:
			resCode = http.StatusUnauthorized
		case custom_error.CodeClientForbidden:
			resCode = http.StatusForbidden
		case custom_error.CodeUnprocessableEntity:
			resCode = http.StatusUnprocessableEntity
		}

		if resCode != http.StatusInternalServerError {
			//TODO: add logger

			if ew.Err != nil {
				return responseJSON(nil, ew.Err.Error(), ew.Validation, resCode)
			}
			return responseJSON(nil, ew.Message, ew.Validation, resCode)
		}

		//TODO: add logger

		return responseJSON(nil, custom_error.MsgServerError, nil, resCode)
	}

	//TODO: add logger

	return responseJSON(nil, custom_error.MsgServerError, nil, resCode)
}

func responseJSON(data interface{}, msg string, errors interface{}, code int) response.JsonResponse {
	return response.JsonResponse{
		Code:    code,
		Message: msg,
		Errors:  errors,
		Data:    data,
	}
}
