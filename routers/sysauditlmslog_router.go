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
		api.POST("CreateLmsLog", middlewares.AuthMiddleware(global.OWEN_DB), middlewares.AuditMiddleware(global.OWEN_DB), Router.sysauditlmslogApi.CreateLmsLog)
		api.DELETE("DeleteLmsLog", middlewares.AuthMiddleware(global.OWEN_DB), middlewares.AuditMiddleware(global.OWEN_DB), Router.sysauditlmslogApi.DeleteLmsLog)
		api.PUT("UpdateLmsLog", middlewares.AuthMiddleware(global.OWEN_DB), middlewares.AuditMiddleware(global.OWEN_DB), Router.sysauditlmslogApi.UpdateLmsLog)
		api.GET("QueryLmsLog", middlewares.AuditMiddleware(global.OWEN_DB), Router.sysauditlmslogApi.QueryLmsLog)
	}
}
