package models

// 角色模型
type Role struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);unique"` // 角色名称
	Permissions string `gorm:"type:text"`                // 权限列表，使用逗号分隔
}

// 用户模型
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
	Role     Role `gorm:"foreignKey:RoleID"`
}
