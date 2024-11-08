package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterUserRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/user")
	{
		// 使用自定义的限流配置（每秒2个请求，最多10个并发）
		// 使用自定义的熔断器配置（最大请求数10，超时3秒）
		api.GET("GetUsers", middlewares.RateLimiter(2, 10), middlewares.CircuitBreaker(10, 3*time.Second), ApiApp.userController.GetUsers)
		api.POST("CreateUser", ApiApp.userController.CreateUser)
	}
}
