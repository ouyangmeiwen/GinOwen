package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterSysauditlmslogRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/Sysauditlmslog")
	{
		api.POST("CreateLmsLog", middlewares.AuthMiddleware(), middlewares.AuditMiddleware(), ApiGroup.sysauditlmslogApi.CreateLmsLog)
		api.DELETE("DeleteLmsLog", middlewares.AuthMiddleware(), middlewares.AuditMiddleware(), ApiGroup.sysauditlmslogApi.DeleteLmsLog)
		api.PUT("UpdateLmsLog", middlewares.AuthMiddleware(), middlewares.AuditMiddleware(), ApiGroup.sysauditlmslogApi.UpdateLmsLog)
		api.GET("QueryLmsLog", middlewares.AuditMiddleware(), ApiGroup.sysauditlmslogApi.QueryLmsLog)
	}
}
