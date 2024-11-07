package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter 限流中间件
// 每个请求可以传递自定义的速率限制，默认每秒1个请求，最多5个并发
func RateLimiter(rateLimit float64, burst int) gin.HandlerFunc {
	// 如果没有传递参数，则使用默认值
	if rateLimit == 0 {
		rateLimit = 1 // 默认每秒1个请求
	}
	if burst == 0 {
		burst = 5 // 默认最多允许5个并发
	}

	// 创建一个新的限流器，rate.Limit 类型是 float64，转换 rateLimit 为 rate.Limit 类型
	limiter := rate.NewLimiter(rate.Limit(rateLimit), burst)

	return func(c *gin.Context) {
		// 检查请求是否符合限流条件
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}
		c.Next()
	}
}
