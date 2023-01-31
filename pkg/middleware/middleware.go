package middleware

import (
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type CustomRateLimiter struct {
	lastRequestTimeMap map[string]time.Time
	requestCountMap    map[string]int
	mu                 sync.Mutex
	limit              int
	window             time.Duration
}

// NewCustomRateLimiter returns a new instance of CustomRateLimiter
func NewCustomRateLimiter(limit int, window time.Duration) *CustomRateLimiter {
	return &CustomRateLimiter{
		lastRequestTimeMap: make(map[string]time.Time),
		requestCountMap:    make(map[string]int),
		limit:              limit,
		window:             window,
	}
}

// Middleware is the custom rate limiting middleware
func (r *CustomRateLimiter) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r.mu.Lock()
		defer r.mu.Unlock()

		identity := c.RealIP()
		lastRequestTime, ok := r.lastRequestTimeMap[identity]
		if !ok {
			r.lastRequestTimeMap[identity] = time.Now()
			r.requestCountMap[identity] = 1
			return next(c)
		}

		if time.Since(lastRequestTime) > r.window {
			r.lastRequestTimeMap[identity] = time.Now()
			r.requestCountMap[identity] = 1
			return next(c)
		}

		requestCount := r.requestCountMap[identity]
		if requestCount < r.limit {
			r.requestCountMap[identity] = requestCount + 1
			return next(c)
		}

		return echo.NewHTTPError(429, "Too Many Requests")
	}
}
