package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func DelayMiddleware(delay time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		time.Sleep(delay)
		ctx.Next()
	}
}
