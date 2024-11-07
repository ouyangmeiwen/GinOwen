package models

import "gorm.io/gorm"

// Order 订单模型
type Order struct {
	gorm.Model
	UserID uint    `json:"user_id"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}
