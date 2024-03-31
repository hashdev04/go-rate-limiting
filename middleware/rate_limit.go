package middleware

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

const limit uint = 2

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

// This makes it so each ip can only make 5 requests per second
var store = ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
	Rate:  time.Minute,
	Limit: limit,
})

var EnsureRateLimit = ratelimit.RateLimiter(store, &ratelimit.Options{
	ErrorHandler: errorHandler,
	KeyFunc:      keyFunc,
})
