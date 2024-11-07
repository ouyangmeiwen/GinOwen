package repositories

import (
	"GINOWEN/global"
	"GINOWEN/models"
)

// GormOrderRepository 使用 GORM 实现 OrderRepository
type GormOrderRepository struct {
}

// CreateOrder 创建订单
func (r *GormOrderRepository) CreateOrder(order *models.Order) error {
	return global.OWEN_DB.Create(order).Error
}

// GetAllOrders 获取所有订单
func (r *GormOrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := global.OWEN_DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
