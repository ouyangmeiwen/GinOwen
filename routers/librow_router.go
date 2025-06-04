package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

type LibRowRouter struct {
}

// RegisterOrderRoutes 注册订单路由
func RegisterLibRowRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.RouterPre + "/LibRow").Use(middlewares.AuditLogMiddleware()).Use(middlewares.AuthMiddleware("librow"))
	{
		api.GET("QueryRows", ApiGroup.librowApi.QueryRows)
	}
}
