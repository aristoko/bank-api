package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-ID"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(RequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		// Simpan ke header response dan context
		c.Writer.Header().Set(RequestIDKey, requestID)
		c.Set(RequestIDKey, requestID)

		c.Next()
	}
}
