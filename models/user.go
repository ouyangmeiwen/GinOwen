package models

// User 用户模型
// @Description 用户数据模型
// @Tags User
type User struct {
	// @ID
	ID uint `json:"id" gorm:"primaryKey"`
	// @Name 用户名
	Name string `json:"name" gorm:"type:varchar(100)" example:"John Doe"`
}
