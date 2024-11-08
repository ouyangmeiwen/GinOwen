package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterLibitemRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/libitem").Use(middlewares.AuthMiddleware())
	{
		api.POST("ImportExcel", ApiGroup.libItemApi.ImportExcel)
	}
}
