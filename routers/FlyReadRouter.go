package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterFlyReadRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.RouterPre + "/FlyRead").Use(middlewares.AuditLogMiddleware()) //.Use(middlewares.AuthMiddleware("flyread"))
	{
		api.GET("Hello", ApiGroup.flyreadApi.Hello)
	}
}
