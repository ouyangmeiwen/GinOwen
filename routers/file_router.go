package routers

import (
	"GINOWEN/global"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterUploadfileRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/file")
	{
		api.POST("UploadFile", ApiGroup.uploadfileApi.UploadFile)
	}
}
