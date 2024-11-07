package services

import (
	"GINOWEN/models"
)

// OrderService 订单服务
type OrderService struct {
}

// GetAllOrders 获取所有订单
func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return RepoApp.orderRepo.GetAllOrders()
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(order *models.Order) error {
	return RepoApp.orderRepo.CreateOrder(order)
}
