package routers

import (
	"GINOWEN/global"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterSysauditlmslogRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/Sysauditlmslog")
	{
		api.POST("CreateLmsLog", Router.sysauditlmslogApi.CreateLmsLog)
		api.DELETE("DeleteLmsLog", Router.sysauditlmslogApi.DeleteLmsLog)
		api.PUT("UpdateLmsLog", Router.sysauditlmslogApi.UpdateLmsLog)
		api.GET("QueryLmsLog", Router.sysauditlmslogApi.QueryLmsLog)
	}
}
