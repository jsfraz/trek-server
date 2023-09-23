package middlewares

import (
	"log"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
)

// Logging middleware.
//
//	@param skip
//	@return gin.HandlerFunc
func Logging(skip []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if !slices.Contains(skip, path) {
			// Stop timer
			now := time.Now()
			latency := now.Sub(start)

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

			if errorMessage == "" {
				log.Printf("| %3d | %13v | %15s | %-7s | %#v",
					statusCode,
					latency,
					clientIP,
					method,
					path,
				)
			} else {
				log.Printf("| %3d | %13v | %15s | %-7s | %#v | %s",
					statusCode,
					latency,
					clientIP,
					method,
					path,
					errorMessage,
				)
			}
		}
	}
}
