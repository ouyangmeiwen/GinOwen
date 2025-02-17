package routers

import (
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterIPRoutes(r *gin.Engine) {
	api := r.Group("IP").Use(middlewares.AuditLogMiddleware()).Use(middlewares.AuthMiddleware("admin"))
	{
		api.POST("AddBlackList", ApiGroup.blacklistApi.AddBlackList)
		api.POST("UnLockIp", ApiGroup.blacklistApi.UnLockIp)
		api.GET("GetBlackList", ApiGroup.blacklistApi.GetBlackList)
	}

	mq := r.Group("MQ").Use(middlewares.AuditLogMiddleware()).Use(middlewares.AuthMiddleware("admin"))
	{
		mq.POST("SendRabbitMQMsg", ApiGroup.blacklistApi.SendRabbitMQMsg)
	}
}
