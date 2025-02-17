package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterSysauditlmslogRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.RouterPre + "/Sysauditlmslog").Use(middlewares.AuditLogMiddleware()).Use(middlewares.AuthMiddleware("sysauditlmslog"))
	{
		api.POST("CreateLmsLog", ApiGroup.sysauditlmslogApi.CreateLmsLog)
		api.DELETE("DeleteLmsLog", ApiGroup.sysauditlmslogApi.DeleteLmsLog)
		api.PUT("UpdateLmsLog", ApiGroup.sysauditlmslogApi.UpdateLmsLog)
		api.GET("QueryLmsLog", ApiGroup.sysauditlmslogApi.QueryLmsLog)
	}
}
