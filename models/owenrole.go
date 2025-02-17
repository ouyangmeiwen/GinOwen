package models

// 角色模型
type OwenRole struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);unique"` // 角色名称
	Permissions string `gorm:"type:text"`                // 权限列表，使用逗号分隔
}

// TableName Libitem's table name
func (*OwenRole) TableName() string {
	return "owen_role"
}
