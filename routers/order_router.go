package routers

import (
	"GINOWEN/global"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterOrderRoutes(r *gin.Engine) {
	api := r.Group(global.OWEN_CONFIG.System.Pre + "/order")
	{
		api.GET("GetOrders", Router.orderController.GetOrders)
		api.POST("CreateOrder", Router.orderController.CreateOrder)
	}
}
