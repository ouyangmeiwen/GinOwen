// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAbpuseraccount = "abpuseraccounts"

// Abpuseraccount mapped from table <abpuseraccounts>
type Abpuseraccount struct {
	ID                   int64      `gorm:"column:Id;type:bigint(20);primaryKey;autoIncrement:true" json:"Id"`
	CreationTime         time.Time  `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID        *int64     `gorm:"column:CreatorUserId;type:bigint(20)" json:"CreatorUserId"`
	LastModificationTime *time.Time `gorm:"column:LastModificationTime;type:datetime(6)" json:"LastModificationTime"`
	LastModifierUserID   *int64     `gorm:"column:LastModifierUserId;type:bigint(20)" json:"LastModifierUserId"`
	IsDeleted            []uint8    `gorm:"column:IsDeleted;type:bit(1);not null" json:"IsDeleted"`
	DeleterUserID        *int64     `gorm:"column:DeleterUserId;type:bigint(20)" json:"DeleterUserId"`
	DeletionTime         *time.Time `gorm:"column:DeletionTime;type:datetime(6)" json:"DeletionTime"`
	TenantID             *int64     `gorm:"column:TenantId;type:int(11)" json:"TenantId"`
	UserID               int64      `gorm:"column:UserId;type:bigint(20);not null" json:"UserId"`
	UserLinkID           *int64     `gorm:"column:UserLinkId;type:bigint(20)" json:"UserLinkId"`
	UserName             *string    `gorm:"column:UserName;type:varchar(256)" json:"UserName"`
	EmailAddress         *string    `gorm:"column:EmailAddress;type:varchar(256)" json:"EmailAddress"`
}

// TableName Abpuseraccount's table name
func (*Abpuseraccount) TableName() string {
	return TableNameAbpuseraccount
}
