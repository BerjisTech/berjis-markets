package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const requestIDHeader = "X-Request-Id"

// RequestID injects a unique ID per request for tracing.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.Header.Get(requestIDHeader)
		if id == "" {
			id = uuid.NewString()
		}
		c.Set(requestIDHeader, id)
		c.Writer.Header().Set(requestIDHeader, id)
		c.Next()
	}
}
