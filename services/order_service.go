package services

import (
	"GINOWEN/models"
	"GINOWEN/repositories"
)

// OrderService 订单服务
type OrderService struct {
	OrderRepo repositories.OrderRepository
}

// NewOrderService 创建 OrderService
func NewOrderService(orderRepo repositories.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

// GetAllOrders 获取所有订单
func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.OrderRepo.GetAllOrders()
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.OrderRepo.CreateOrder(order)
}
