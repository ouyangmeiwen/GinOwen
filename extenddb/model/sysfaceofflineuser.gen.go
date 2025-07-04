// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysfaceofflineuser = "sysfaceofflineuser"

// Sysfaceofflineuser mapped from table <sysfaceofflineuser>
type Sysfaceofflineuser struct {
	ID                   string     `gorm:"column:Id;type:varchar(32);primaryKey" json:"Id"`
	CreationTime         time.Time  `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID        *int64     `gorm:"column:CreatorUserId;type:bigint" json:"CreatorUserId"`
	LastModificationTime *time.Time `gorm:"column:LastModificationTime;type:datetime(6)" json:"LastModificationTime"`
	LastModifierUserID   *int64     `gorm:"column:LastModifierUserId;type:bigint" json:"LastModifierUserId"`
	IsDeleted            []uint8    `gorm:"column:IsDeleted;type:bit(1);not null" json:"IsDeleted"`
	DeleterUserID        *int64     `gorm:"column:DeleterUserId;type:bigint" json:"DeleterUserId"`
	DeletionTime         *time.Time `gorm:"column:DeletionTime;type:datetime(6)" json:"DeletionTime"`
	GroupID              *string    `gorm:"column:GroupId;type:varchar(32)" json:"GroupId"`
	UserID               *string    `gorm:"column:UserId;type:varchar(32)" json:"UserId"`
	UserInfo             *string    `gorm:"column:UserInfo;type:varchar(256)" json:"UserInfo"`
	Image                *string    `gorm:"column:Image;type:longtext" json:"Image"`
	Remark               *string    `gorm:"column:Remark;type:varchar(256)" json:"Remark"`
}

// TableName Sysfaceofflineuser's table name
func (*Sysfaceofflineuser) TableName() string {
	return TableNameSysfaceofflineuser
}
