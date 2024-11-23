package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterLibitemRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/Libitems").Use(middlewares.AuthMiddleware("libitem"))
	{
		api.POST("ImportExcelByName", ApiGroup.libItemApi.ImportExcelByName)
		api.POST("ImportExcelByIndex", ApiGroup.libItemApi.ImportExcelByIndex)
	}
}
