package routers

import (
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

type IPRouter struct {
}

func RegisterIPRoutes(r *gin.Engine) {
	api := r.Group("IP").Use(middlewares.AuditLogMiddleware()).Use(middlewares.AuthMiddleware("admin"))
	{
		api.POST("AddBlackList", ApiGroup.ipApi.AddBlackList)
		api.POST("UnLockIp", ApiGroup.ipApi.UnLockIp)
	}
}
