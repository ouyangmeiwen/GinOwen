package models

// TestOrder 订单模型
// @Description 订单数据模型
// @Tags TestOrder
type TestOrder struct {
	// @ID
	ID       uint    `json:"id" gorm:"primaryKey"`
	Amount   float64 `json:"amount" example:"100.50"`
	Customer string  `json:"customer" example:"John Doe"`
}

// TableName Sysauditlmslog's table name
func (*TestOrder) TableName() string {
	return "testorder"
}
