package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
)

// OrderService 订单服务
type OrderService struct {
}

// GetAllOrders 获取所有订单
func (s *OrderService) GetAllOrders() ([]models.TestOrder, error) {
	var orders []models.TestOrder
	if err := global.OWEN_DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(order *models.TestOrder) error {
	return global.OWEN_DB.Create(order).Error
}
