package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiterConfig controls limiter behavior.
type RateLimiterConfig struct {
	RequestsPerSecond rate.Limit
	Burst             int
}

type clientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// RateLimiter enforces a token bucket per client IP.
func RateLimiter(cfg RateLimiterConfig) gin.HandlerFunc {
	var (
		mu       sync.Mutex
		visitors = make(map[string]*clientLimiter)
	)

	getLimiter := func(ip string) *rate.Limiter {
		mu.Lock()
		defer mu.Unlock()

		v, exists := visitors[ip]
		if !exists {
			limiter := rate.NewLimiter(cfg.RequestsPerSecond, cfg.Burst)
			visitors[ip] = &clientLimiter{limiter: limiter, lastSeen: time.Now()}
			return limiter
		}
		v.lastSeen = time.Now()
		return v.limiter
	}

	// Cleanup goroutine.
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			mu.Lock()
			for ip, v := range visitors {
				if time.Since(v.lastSeen) > 10*time.Minute {
					delete(visitors, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		limiter := getLimiter(c.ClientIP())
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}
		c.Next()
	}
}
