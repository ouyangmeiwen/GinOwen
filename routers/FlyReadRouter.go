package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterFlyReadRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.RouterPre + "/FlyRead").Use(middlewares.AuditLogMiddleware()).Use(middlewares.AuthMiddleware("flyread"))
	{
		api.GET("Hello", ApiGroup.flyreadApi.Hello)
		api.GET("GetFlyReadSetting", ApiGroup.flyreadApi.GetFlyReadSetting)
		api.POST("SetFlyReadSetting", ApiGroup.flyreadApi.SetFlyReadSetting)
		api.GET("GetFlyReadToken", ApiGroup.flyreadApi.GetFlyReadToken)
		api.POST("UploadLibItem", ApiGroup.flyreadApi.UploadLibItem)
		api.POST("UploadTenant", ApiGroup.flyreadApi.UploadTenant)
		api.POST("UploadStruct", ApiGroup.flyreadApi.UploadStruct)
		api.POST("UploadLibItemLoc", ApiGroup.flyreadApi.UploadLibItemLoc)
		api.POST("UploadRow", ApiGroup.flyreadApi.UploadRow)
		api.POST("Inventory", ApiGroup.flyreadApi.Inventory)
		api.GET("GetRobotRouterlist", ApiGroup.flyreadApi.GetRobotRouterlist)
		api.POST("InventoryHis", ApiGroup.flyreadApi.InventoryHis)
		api.POST("InventoryList", ApiGroup.flyreadApi.InventoryList)
	}
}
