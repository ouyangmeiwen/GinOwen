package routers

import (
	"GINOWEN/global"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterSysauditlmslogRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/Sysauditlmslog")
	{
		api.POST("CreateLmsLog", Router.sysauditlmslogController.CreateLmsLog)
		api.DELETE("DeleteLmsLog", Router.sysauditlmslogController.DeleteLmsLog)
		api.PUT("UpdateLmsLog", Router.sysauditlmslogController.UpdateLmsLog)
		api.GET("QueryLmsLog", Router.sysauditlmslogController.QueryLmsLog)
	}
}
