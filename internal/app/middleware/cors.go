package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware is a middleware that handles CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// CustomCORSMiddleware is a middleware that handles CORS with custom configuration
func CustomCORSMiddleware(origins []string, methods []string, headers []string, maxAge time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set allowed origins
		origin := c.Request.Header.Get("Origin")
		for _, allowedOrigin := range origins {
			if origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		// Set allowed methods
		if len(methods) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Methods", methods[0])
			for i := 1; i < len(methods); i++ {
				c.Writer.Header().Add("Access-Control-Allow-Methods", methods[i])
			}
		}

		// Set allowed headers
		if len(headers) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Headers", headers[0])
			for i := 1; i < len(headers); i++ {
				c.Writer.Header().Add("Access-Control-Allow-Headers", headers[i])
			}
		}

		// Set max age
		if maxAge > 0 {
			c.Writer.Header().Set("Access-Control-Max-Age", maxAge.String())
		}

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
