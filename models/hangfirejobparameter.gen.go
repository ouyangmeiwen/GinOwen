// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameHangfirejobparameter = "hangfirejobparameter"

// Hangfirejobparameter mapped from table <hangfirejobparameter>
type Hangfirejobparameter struct {
	ID    int64   `gorm:"column:Id;type:int(11);primaryKey;autoIncrement:true" json:"Id"`
	JobID int64   `gorm:"column:JobId;type:int(11);not null" json:"JobId"`
	Name  string  `gorm:"column:Name;type:varchar(40);not null" json:"Name"`
	Value *string `gorm:"column:Value;type:longtext" json:"Value"`
}

// TableName Hangfirejobparameter's table name
func (*Hangfirejobparameter) TableName() string {
	return TableNameHangfirejobparameter
}
