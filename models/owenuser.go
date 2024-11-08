package models

// 用户模型
type OwenUser struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
	Role     OwenRole `gorm:"foreignKey:RoleID"`
}

// TableName Libitem's table name
func (*OwenUser) TableName() string {
	return "owen_user"
}
