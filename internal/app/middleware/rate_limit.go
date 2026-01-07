package middleware

import (
	"sync"
	"time"

	"github.com/captainkie/websync-api/pkg/errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiterConfig holds configuration for rate limiting
type RateLimiterConfig struct {
	RequestsPerMinute int
	BurstSize         int
	BlockDuration     time.Duration
}

// DefaultRateLimiterConfig returns default rate limiter configuration
func DefaultRateLimiterConfig() RateLimiterConfig {
	return RateLimiterConfig{
		RequestsPerMinute: 100,
		BurstSize:         50,
		BlockDuration:     5 * time.Minute,
	}
}

// IPRateLimiterManager stores rate limiters for different IPs
type IPRateLimiterManager struct {
	ips    map[string]*rate.Limiter
	mu     *sync.RWMutex
	config RateLimiterConfig
}

// NewIPRateLimiterManager creates a new IP rate limiter manager
func NewIPRateLimiterManager(config RateLimiterConfig) *IPRateLimiterManager {
	return &IPRateLimiterManager{
		ips:    make(map[string]*rate.Limiter),
		mu:     &sync.RWMutex{},
		config: config,
	}
}

// GetLimiter returns the rate limiter for the provided IP address
func (i *IPRateLimiterManager) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(i.config.RequestsPerMinute/60), i.config.BurstSize)
		i.ips[ip] = limiter
	}

	return limiter
}

// Cleanup removes old IP entries
func (i *IPRateLimiterManager) Cleanup() {
	i.mu.Lock()
	defer i.mu.Unlock()

	for ip := range i.ips {
		delete(i.ips, ip)
	}
}

// GlobalRateLimiter stores the global rate limiter
var globalLimiter = rate.NewLimiter(rate.Limit(1000/60), 500) // 1000 requests per minute with burst of 500

// ipRateLimiterManager stores rate limiters for different IPs
var ipRateLimiterManager = NewIPRateLimiterManager(DefaultRateLimiterConfig())

// RateLimiter is a middleware that limits the rate of requests
func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !globalLimiter.Allow() {
			err := errors.NewRateLimitError(
				"Global rate limit exceeded. Please try again later.",
				int(ipRateLimiterManager.config.BlockDuration.Seconds()),
			)
			c.Error(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

// IPRateLimiter is a middleware that limits the rate of requests per IP
func IPRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := ipRateLimiterManager.GetLimiter(ip)

		if !limiter.Allow() {
			err := errors.NewRateLimitError(
				"Rate limit exceeded for your IP. Please try again later.",
				int(ipRateLimiterManager.config.BlockDuration.Seconds()),
			)
			c.Error(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

// RouteRateLimiter creates a rate limiter for specific routes
func RouteRateLimiter(requestsPerMinute, burstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(requestsPerMinute/60), burstSize)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			err := errors.NewRateLimitError(
				"Rate limit exceeded for this endpoint. Please try again later.",
				int(ipRateLimiterManager.config.BlockDuration.Seconds()),
			)
			c.Error(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

// StartCleanupRoutine starts a routine to clean up old IP entries
func StartCleanupRoutine() {
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			ipRateLimiterManager.Cleanup()
		}
	}()
}
