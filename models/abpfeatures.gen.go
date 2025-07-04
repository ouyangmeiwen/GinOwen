// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameAbpfeature = "abpfeatures"

// Abpfeature mapped from table <abpfeatures>
type Abpfeature struct {
	ID            int64     `gorm:"column:Id;type:bigint;primaryKey;autoIncrement:true" json:"Id"`
	CreationTime  time.Time `gorm:"column:CreationTime;type:datetime(6);not null" json:"CreationTime"`
	CreatorUserID *int64    `gorm:"column:CreatorUserId;type:bigint" json:"CreatorUserId"`
	TenantID      *int64    `gorm:"column:TenantId;type:int" json:"TenantId"`
	Name          string    `gorm:"column:Name;type:varchar(128);not null" json:"Name"`
	Value         string    `gorm:"column:Value;type:varchar(2000);not null" json:"Value"`
	Discriminator string    `gorm:"column:Discriminator;type:longtext;not null" json:"Discriminator"`
	EditionID     *int64    `gorm:"column:EditionId;type:int" json:"EditionId"`
}

// TableName Abpfeature's table name
func (*Abpfeature) TableName() string {
	return TableNameAbpfeature
}
