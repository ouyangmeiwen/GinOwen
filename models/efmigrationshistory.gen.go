// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameEfmigrationshistory = "__efmigrationshistory"

// Efmigrationshistory mapped from table <__efmigrationshistory>
type Efmigrationshistory struct {
	MigrationID    string `gorm:"column:MigrationId;type:varchar(95);primaryKey" json:"MigrationId"`
	ProductVersion string `gorm:"column:ProductVersion;type:varchar(32);not null" json:"ProductVersion"`
}

// TableName Efmigrationshistory's table name
func (*Efmigrationshistory) TableName() string {
	return TableNameEfmigrationshistory
}
