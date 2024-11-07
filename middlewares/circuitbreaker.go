package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

var cbSettings = gobreaker.Settings{
	Name:        "HTTP Circuit Breaker",
	MaxRequests: 5,
	Interval:    0,
	Timeout:     5 * time.Second,
}
var breaker = gobreaker.NewCircuitBreaker(cbSettings)

// CircuitBreaker 熔断中间件
func CircuitBreaker() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := breaker.Execute(func() (interface{}, error) {
			c.Next()
			return nil, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"error": "Service unavailable",
			})
		}
	}
}
