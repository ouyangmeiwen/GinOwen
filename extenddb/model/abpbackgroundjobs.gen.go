// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAbpbackgroundjob = "abpbackgroundjobs"

// Abpbackgroundjob mapped from table <abpbackgroundjobs>
type Abpbackgroundjob struct {
	ID            int64      `gorm:"column:Id;type:bigint;primaryKey;autoIncrement:true" json:"Id"`
	CreationTime  time.Time  `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID *int64     `gorm:"column:CreatorUserId;type:bigint" json:"CreatorUserId"`
	JobType       string     `gorm:"column:JobType;type:varchar(512);not null" json:"JobType"`
	JobArgs       string     `gorm:"column:JobArgs;type:longtext;not null" json:"JobArgs"`
	TryCount      int64      `gorm:"column:TryCount;type:smallint;not null" json:"TryCount"`
	NextTryTime   time.Time  `gorm:"column:NextTryTime;type:datetime(6);not null" json:"NextTryTime"`
	LastTryTime   *time.Time `gorm:"column:LastTryTime;type:datetime(6)" json:"LastTryTime"`
	IsAbandoned   []uint8    `gorm:"column:IsAbandoned;type:bit(1);not null" json:"IsAbandoned"`
	Priority      int64      `gorm:"column:Priority;type:tinyint unsigned;not null" json:"Priority"`
}

// TableName Abpbackgroundjob's table name
func (*Abpbackgroundjob) TableName() string {
	return TableNameAbpbackgroundjob
}
