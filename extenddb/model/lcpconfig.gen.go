// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLcpconfig = "lcpconfig"

// Lcpconfig mapped from table <lcpconfig>
type Lcpconfig struct {
	ID                   string     `gorm:"column:Id;type:varchar(32);primaryKey" json:"Id"`
	CreationTime         time.Time  `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID        *int64     `gorm:"column:CreatorUserId;type:bigint" json:"CreatorUserId"`
	LastModificationTime *time.Time `gorm:"column:LastModificationTime;type:datetime(6)" json:"LastModificationTime"`
	LastModifierUserID   *int64     `gorm:"column:LastModifierUserId;type:bigint" json:"LastModifierUserId"`
	IsDeleted            []uint8    `gorm:"column:IsDeleted;type:bit(1);not null" json:"IsDeleted"`
	DeleterUserID        *int64     `gorm:"column:DeleterUserId;type:bigint" json:"DeleterUserId"`
	DeletionTime         *time.Time `gorm:"column:DeletionTime;type:datetime(6)" json:"DeletionTime"`
	TargetID             string     `gorm:"column:TargetId;type:varchar(32);not null" json:"TargetId"`
	TargetCode           *string    `gorm:"column:TargetCode;type:varchar(32)" json:"TargetCode"`
	TargetName           *string    `gorm:"column:TargetName;type:varchar(128)" json:"TargetName"`
	Directory            *string    `gorm:"column:Directory;type:varchar(256)" json:"Directory"`
	FileName             string     `gorm:"column:FileName;type:varchar(64);not null" json:"FileName"`
	Content              *string    `gorm:"column:Content;type:longtext" json:"Content"`
	IsLost               []uint8    `gorm:"column:IsLost;type:bit(1);not null" json:"IsLost"`
	Remark               *string    `gorm:"column:Remark;type:varchar(256)" json:"Remark"`
	TenantID             int64      `gorm:"column:TenantId;type:int;not null" json:"TenantId"`
}

// TableName Lcpconfig's table name
func (*Lcpconfig) TableName() string {
	return TableNameLcpconfig
}
