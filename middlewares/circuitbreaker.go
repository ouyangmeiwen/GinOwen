package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

// CircuitBreaker 熔断中间件
// 可以通过传递自定义的参数来控制熔断器的行为
func CircuitBreaker(maxRequests uint32, timeout time.Duration) gin.HandlerFunc {
	// 默认配置
	if maxRequests == 0 {
		maxRequests = 5 // 默认最大请求数
	}
	if timeout == 0 {
		timeout = 5 * time.Second // 默认熔断超时时间
	}

	// 创建一个新的熔断器
	cbSettings := gobreaker.Settings{
		Name:        "HTTP Circuit Breaker",
		MaxRequests: maxRequests,
		Interval:    0, // 可选配置：用来控制熔断器的状态更新频率，通常设为0
		Timeout:     timeout,
	}
	breaker := gobreaker.NewCircuitBreaker(cbSettings)

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
