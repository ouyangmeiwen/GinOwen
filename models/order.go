package models

// Order 订单模型
// @Description 订单数据模型
// @Tags Order
type Order struct {
	// @ID
	ID       uint    `json:"id" gorm:"primaryKey"`
	Amount   float64 `json:"amount" example:"100.50"`
	Customer string  `json:"customer" example:"John Doe"`
}
