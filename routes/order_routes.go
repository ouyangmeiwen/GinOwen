package routes

import (
	"GINOWEN/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes 注册订单路由
func RegisterOrderRoutes(r *gin.Engine, orderController *controllers.OrderController) {
	api := r.Group("/api/orders")
	{
		api.GET("/", orderController.GetOrders)
		api.POST("/", orderController.CreateOrder)
	}
}
