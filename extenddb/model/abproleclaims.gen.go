// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAbproleclaim = "abproleclaims"

// Abproleclaim mapped from table <abproleclaims>
type Abproleclaim struct {
	ID            int64     `gorm:"column:Id;type:bigint;primaryKey;autoIncrement:true" json:"Id"`
	CreationTime  time.Time `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID *int64    `gorm:"column:CreatorUserId;type:bigint" json:"CreatorUserId"`
	TenantID      *int64    `gorm:"column:TenantId;type:int" json:"TenantId"`
	RoleID        int64     `gorm:"column:RoleId;type:int;not null" json:"RoleId"`
	ClaimType     *string   `gorm:"column:ClaimType;type:varchar(256)" json:"ClaimType"`
	ClaimValue    *string   `gorm:"column:ClaimValue;type:longtext" json:"ClaimValue"`
}

// TableName Abproleclaim's table name
func (*Abproleclaim) TableName() string {
	return TableNameAbproleclaim
}
