package routers

import (
	"GINOWEN/global"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterLibitemRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/libitem")
	{
		api.POST("ImportExcel", ApiApp.libItemController.ImportExcel)
	}
}
