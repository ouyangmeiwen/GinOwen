package models

// 用户模型
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
	Role     Role `gorm:"foreignKey:RoleID"`
}

// TableName Libitem's table name
func (*User) TableName() string {
	return "owen_user"
}
