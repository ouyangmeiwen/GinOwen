// initrouter/order.go
package initcontrollers

import (
	"GINOWEN/controllers"
	"GINOWEN/repositories"
	"GINOWEN/services"

	"gorm.io/gorm"
)

// InitOrderController 初始化订单控制器
func InitOrderController(db *gorm.DB) *controllers.OrderController {
	orderRepo := &repositories.GormOrderRepository{DB: db}
	orderService := services.NewOrderService(orderRepo)
	return controllers.NewOrderController(orderService)
}
