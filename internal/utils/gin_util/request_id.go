package gin_util

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
)

func RequestIDMiddleware(c *gin.Context) {
	requestID := c.GetHeader("X-Request-ID")

	if requestID == "" {
		requestID, _ = uuid.GenerateUUID()
		c.Request.Header.Set("X-Request-ID", requestID)
	}

	c.Header("X-Request-ID", requestID)
}
