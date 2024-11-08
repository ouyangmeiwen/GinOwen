package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterSysauditlmslogRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/Sysauditlmslog").Use(middlewares.AuditMiddleware())
	{
		api.POST("CreateLmsLog", middlewares.AuthMiddleware("sysauditlmslog_create"), ApiGroup.sysauditlmslogApi.CreateLmsLog)
		api.DELETE("DeleteLmsLog", middlewares.AuthMiddleware("sysauditlmslog_delete"), ApiGroup.sysauditlmslogApi.DeleteLmsLog)
		api.PUT("UpdateLmsLog", middlewares.AuthMiddleware("sysauditlmslog_update"), ApiGroup.sysauditlmslogApi.UpdateLmsLog)
		api.GET("QueryLmsLog", ApiGroup.sysauditlmslogApi.QueryLmsLog)
	}
}
