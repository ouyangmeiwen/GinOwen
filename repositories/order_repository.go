package repositories

import (
	"GINOWEN/models"

	"gorm.io/gorm"
)

// OrderRepository 定义订单仓库接口
type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetAllOrders() ([]models.Order, error)
}

// GormOrderRepository 使用 GORM 实现 OrderRepository
type GormOrderRepository struct {
	DB *gorm.DB
}

// CreateOrder 创建订单
func (r *GormOrderRepository) CreateOrder(order *models.Order) error {
	return r.DB.Create(order).Error
}

// GetAllOrders 获取所有订单
func (r *GormOrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := r.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
