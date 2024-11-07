package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterOrderRoutes(r *gin.Engine) {
	api := r.Group("api/orders")
	{
		api.GET("GetOrders", ApiApp.orderController.GetOrders)
		api.POST("CreateOrder", ApiApp.orderController.CreateOrder)
	}
}
