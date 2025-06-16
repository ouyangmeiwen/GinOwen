package models

import (
	"time"
)

const TableNameSystenantconfig = "systenantconfig"

// Systenantconfig mapped from table <systenantconfig>
type Systenantconfig struct {
	ID                   string     `gorm:"column:Id;type:varchar(32);primaryKey" json:"Id"`
	CreationTime         time.Time  `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID        *int64     `gorm:"column:CreatorUserId;type:bigint(20)" json:"CreatorUserId"`
	LastModificationTime *time.Time `gorm:"column:LastModificationTime;type:datetime(6)" json:"LastModificationTime"`
	LastModifierUserID   *int64     `gorm:"column:LastModifierUserId;type:bigint(20)" json:"LastModifierUserId"`
	System               *string    `gorm:"column:System;type:longtext" json:"System"`
	BusinessRule         *string    `gorm:"column:BusinessRule;type:longtext" json:"BusinessRule"`
	Protocol             *string    `gorm:"column:Protocol;type:longtext" json:"Protocol"`
	WeChat               *string    `gorm:"column:WeChat;type:longtext" json:"WeChat"`
	Alipay               *string    `gorm:"column:Alipay;type:longtext" json:"Alipay"`
	Baidu                *string    `gorm:"column:Baidu;type:longtext" json:"Baidu"`
	QQWeiXiao            *string    `gorm:"column:QQWeiXiao;type:longtext" json:"QQWeiXiao"`
	Stat                 *string    `gorm:"column:Stat;type:longtext" json:"Stat"`
	Task                 *string    `gorm:"column:Task;type:longtext" json:"Task"`
	Points               *string    `gorm:"column:Points;type:longtext" json:"Points"`
	Remark               *string    `gorm:"column:Remark;type:varchar(256)" json:"Remark"`
	TenantID             int64      `gorm:"column:TenantId;type:int(11);not null" json:"TenantId"`
}

// TableName Systenantconfig's table name
func (*Systenantconfig) TableName() string {
	return TableNameSystenantconfig
}
