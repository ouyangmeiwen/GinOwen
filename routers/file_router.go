package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterUploadfileRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/file").Use(middlewares.AuthMiddleware("file"))
	{
		api.POST("UploadFile", ApiGroup.uploadfileApi.UploadFile)
		api.GET("DownloadFile", ApiGroup.uploadfileApi.DownloadFile)
	}
}
