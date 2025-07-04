// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameLibpatron = "libpatron"

// Libpatron mapped from table <libpatron>
type Libpatron struct {
	ID                   string     `gorm:"column:Id;type:varchar(32);primaryKey" json:"Id"`
	CreationTime         time.Time  `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID        *int64     `gorm:"column:CreatorUserId;type:bigint" json:"CreatorUserId"`
	LastModificationTime *time.Time `gorm:"column:LastModificationTime;type:datetime(6)" json:"LastModificationTime"`
	LastModifierUserID   *int64     `gorm:"column:LastModifierUserId;type:bigint" json:"LastModifierUserId"`
	IsDeleted            []uint8    `gorm:"column:IsDeleted;type:bit(1);not null" json:"IsDeleted"`
	DeleterUserID        *int64     `gorm:"column:DeleterUserId;type:bigint" json:"DeleterUserId"`
	DeletionTime         *time.Time `gorm:"column:DeletionTime;type:datetime(6)" json:"DeletionTime"`
	Name                 string     `gorm:"column:Name;type:varchar(64);not null" json:"Name"`
	IDCard               *string    `gorm:"column:IdCard;type:varchar(32)" json:"IdCard"`
	Barcode              string     `gorm:"column:Barcode;type:varchar(64);not null" json:"Barcode"`
	IsEnable             []uint8    `gorm:"column:IsEnable;type:bit(1);not null" json:"IsEnable"`
	Sex                  int64      `gorm:"column:Sex;type:tinyint unsigned;not null" json:"Sex"`
	Birthday             *time.Time `gorm:"column:Birthday;type:datetime(6)" json:"Birthday"`
	Password             *string    `gorm:"column:Password;type:varchar(128)" json:"Password"`
	CardTypeID           *string    `gorm:"column:CardTypeId;type:varchar(32)" json:"CardTypeId"`
	CardTypeName         *string    `gorm:"column:CardTypeName;type:varchar(128)" json:"CardTypeName"`
	Tid                  *string    `gorm:"column:Tid;type:varchar(32)" json:"Tid"`
	Points               int64      `gorm:"column:Points;type:int;not null" json:"Points"`
	DepositMoney         int64      `gorm:"column:DepositMoney;type:int;not null" json:"DepositMoney"`
	Balance              int64      `gorm:"column:Balance;type:int;not null" json:"Balance"`
	LateFee              int64      `gorm:"column:LateFee;type:int;not null" json:"LateFee"`
	Phone                *string    `gorm:"column:Phone;type:varchar(32)" json:"Phone"`
	Email                *string    `gorm:"column:Email;type:varchar(64)" json:"Email"`
	Address              *string    `gorm:"column:Address;type:varchar(128)" json:"Address"`
	DepartmentID         *string    `gorm:"column:DepartmentId;type:varchar(32)" json:"DepartmentId"`
	DepartmentName       *string    `gorm:"column:DepartmentName;type:varchar(256)" json:"DepartmentName"`
	ExpireTime           *time.Time `gorm:"column:ExpireTime;type:datetime(6)" json:"ExpireTime"`
	Remark               *string    `gorm:"column:Remark;type:varchar(256)" json:"Remark"`
	OriginType           int64      `gorm:"column:OriginType;type:tinyint unsigned;not null" json:"OriginType"`
	CreateType           int64      `gorm:"column:CreateType;type:tinyint unsigned;not null" json:"CreateType"`
	TenantID             int64      `gorm:"column:TenantId;type:int;not null" json:"TenantId"`
}

// TableName Libpatron's table name
func (*Libpatron) TableName() string {
	return TableNameLibpatron
}
