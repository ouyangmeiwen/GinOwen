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
		api.POST("SetBussiness", ApiGroup.flyreadApi.SetBussiness)
		api.GET("GetInventorySet", ApiGroup.flyreadApi.GetInventorySet)
		api.GET("GetEnableRow", ApiGroup.flyreadApi.GetEnableRow)
		api.GET("GetRobotlist", ApiGroup.flyreadApi.GetRobotlist)
		api.POST("GetCaseCodeImage", ApiGroup.flyreadApi.GetCaseCodeImage)
		api.GET("GetOcrImgs", ApiGroup.flyreadApi.GetOcrImgs)
		api.POST("InventorySet", ApiGroup.flyreadApi.InventorySet)
		api.POST("CreatWork", ApiGroup.flyreadApi.CreatWork)
		api.POST("UpdateWork", ApiGroup.flyreadApi.UpdateWork)
		api.POST("WorkList", ApiGroup.flyreadApi.WorkList)
		api.GET("DeleteWork", ApiGroup.flyreadApi.DeleteWork)
		api.POST("DetailList", ApiGroup.flyreadApi.DetailList)
		api.POST("DetailStatusList", ApiGroup.flyreadApi.DetailStatusList)
		api.POST("InventoryMonthList", ApiGroup.flyreadApi.InventoryMonthList)
		api.POST("BooksNewIndex", ApiGroup.flyreadApi.BooksNewIndex)
		api.GET("GetNotHitRank", ApiGroup.flyreadApi.GetNotHitRank)
		api.GET("GetFaultRank", ApiGroup.flyreadApi.GetFaultRank)
		api.POST("BooksIndex", ApiGroup.flyreadApi.BooksIndex)
		api.POST("BookRankIndex", ApiGroup.flyreadApi.BookRankIndex)
		api.POST("InventoryFlyReadIndex", ApiGroup.flyreadApi.InventoryFlyReadIndex)
		api.GET("GetStructTreeList", ApiGroup.flyreadApi.GetStructTreeList)
		api.GET("GetEnableStruct", ApiGroup.flyreadApi.GetEnableStruct)
	}
}
