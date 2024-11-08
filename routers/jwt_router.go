package routers

import (
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

type JWTRouter struct {
}

// RegisterOrderRoutes 注册订单路由
func RegisterJWTRoutes(r *gin.Engine) {
	api := r.Group("auth").Use(middlewares.AuditMiddleware())
	{
		api.POST("Login", ApiGroup.jwtApi.Login)
		api.POST("Register", middlewares.AuthMiddleware(), ApiGroup.jwtApi.Register)
	}
}
