// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameDascirculatecount = "dascirculatecount"

// Dascirculatecount mapped from table <dascirculatecount>
type Dascirculatecount struct {
	ID            string    `gorm:"column:Id;type:char(36);primaryKey" json:"Id"`
	CreationTime  time.Time `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID *int64    `gorm:"column:CreatorUserId;type:bigint" json:"CreatorUserId"`
	TerminalID    *string   `gorm:"column:TerminalId;type:varchar(32)" json:"TerminalId"`
	TerminalCode  *string   `gorm:"column:TerminalCode;type:varchar(32)" json:"TerminalCode"`
	TerminalName  *string   `gorm:"column:TerminalName;type:varchar(128)" json:"TerminalName"`
	StartTime     time.Time `gorm:"column:StartTime;type:datetime(6);not null" json:"StartTime"`
	EndTime       time.Time `gorm:"column:EndTime;type:datetime(6);not null" json:"EndTime"`
	Type          int64     `gorm:"column:Type;type:tinyint unsigned;not null" json:"Type"`
	Result        int64     `gorm:"column:Result;type:tinyint unsigned;not null" json:"Result"`
	Count         int64     `gorm:"column:Count;type:int;not null" json:"Count"`
	TenantID      int64     `gorm:"column:TenantId;type:int;not null" json:"TenantId"`
}

// TableName Dascirculatecount's table name
func (*Dascirculatecount) TableName() string {
	return TableNameDascirculatecount
}
