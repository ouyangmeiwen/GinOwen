// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysdataitem = "sysdataitem"

// Sysdataitem mapped from table <sysdataitem>
type Sysdataitem struct {
	ID                   string     `gorm:"column:Id;type:varchar(32);primaryKey" json:"Id"`
	CreationTime         time.Time  `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID        *int64     `gorm:"column:CreatorUserId;type:bigint" json:"CreatorUserId"`
	LastModificationTime *time.Time `gorm:"column:LastModificationTime;type:datetime(6)" json:"LastModificationTime"`
	LastModifierUserID   *int64     `gorm:"column:LastModifierUserId;type:bigint" json:"LastModifierUserId"`
	IsDeleted            []uint8    `gorm:"column:IsDeleted;type:bit(1);not null" json:"IsDeleted"`
	DeleterUserID        *int64     `gorm:"column:DeleterUserId;type:bigint" json:"DeleterUserId"`
	DeletionTime         *time.Time `gorm:"column:DeletionTime;type:datetime(6)" json:"DeletionTime"`
	Code                 string     `gorm:"column:Code;type:varchar(32);not null" json:"Code"`
	Name                 string     `gorm:"column:Name;type:varchar(64);not null" json:"Name"`
	IsTree               []uint8    `gorm:"column:IsTree;type:bit(1);not null" json:"IsTree"`
	SortCode             int64      `gorm:"column:SortCode;type:int;not null" json:"SortCode"`
	IsEnable             []uint8    `gorm:"column:IsEnable;type:bit(1);not null" json:"IsEnable"`
	Remark               *string    `gorm:"column:Remark;type:varchar(256)" json:"Remark"`
	ParentID             *string    `gorm:"column:ParentId;type:varchar(32)" json:"ParentId"`
}

// TableName Sysdataitem's table name
func (*Sysdataitem) TableName() string {
	return TableNameSysdataitem
}
