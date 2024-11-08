package routers

import (
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

type JWTRouter struct {
}

// RegisterOrderRoutes 注册订单路由
func RegisterJWTRoutes(r *gin.Engine) {
	api := r.Group("auth")
	{
		api.POST("Login", middlewares.AuditMiddleware(), Router.jwtApi.Login)
		api.POST("Register", middlewares.AuditMiddleware(), Router.jwtApi.Register)
	}
}
