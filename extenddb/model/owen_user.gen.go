// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOwenUser = "owen_user"

// OwenUser mapped from table <owen_user>
type OwenUser struct {
	ID       int64   `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Username *string `gorm:"column:username;type:varchar(255)" json:"username"`
	Password *string `gorm:"column:password;type:varchar(255)" json:"password"`
	RoleID   *int64  `gorm:"column:role_id;type:bigint unsigned" json:"role_id"`
}

// TableName OwenUser's table name
func (*OwenUser) TableName() string {
	return TableNameOwenUser
}
