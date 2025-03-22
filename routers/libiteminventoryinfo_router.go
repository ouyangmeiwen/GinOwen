package routers

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

type LibiteminventoryinfoRouter struct {
}

// RegisterOrderRoutes 注册订单路由
func RegisterInventoryRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.RouterPre + "/Libiteminventory").Use(middlewares.AuditLogMiddleware()).Use(middlewares.AuthMiddleware("libiteminventory"))
	{
		api.POST("CreateInventory", ApiGroup.inventoryApi.CreateInventory)
		api.DELETE("DeleteInventory", ApiGroup.inventoryApi.DeleteInventory)
		api.PUT("UpdateInventory", ApiGroup.inventoryApi.UpdateInventory)
		api.GET("QueryInventory", ApiGroup.inventoryApi.QueryInventory)
	}
}
