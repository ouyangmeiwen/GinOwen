package routers

import (
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterJWTRoutes(r *gin.Engine) {
	api := r.Group("auth").Use(middlewares.AuditLogMiddleware())
	{
		api.POST("Login", ApiGroup.loginApi.Login)
		api.POST("Register", middlewares.AuthMiddleware(), ApiGroup.loginApi.Register)
		api.POST("LoginOut", middlewares.AuthMiddleware(), ApiGroup.loginApi.LoginOut)
		//api.GET("DebugIn", ApiGroup.jwtApi.DebugIn)
		//api.GET("DebugOut", ApiGroup.jwtApi.DebugOut)
	}
}
