// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRescipinfo = "rescipinfo"

// Rescipinfo mapped from table <rescipinfo>
type Rescipinfo struct {
	ID          int64   `gorm:"column:Id;type:int;primaryKey;autoIncrement:true" json:"Id"`
	CIP         string  `gorm:"column:CIP;type:varchar(32);not null" json:"CIP"`
	ISBN        *string `gorm:"column:ISBN;type:varchar(32)" json:"ISBN"`
	Title       string  `gorm:"column:Title;type:varchar(256);not null" json:"Title"`
	Author      *string `gorm:"column:Author;type:varchar(256)" json:"Author"`
	Series      *string `gorm:"column:Series;type:varchar(128)" json:"Series"`
	Publisher   *string `gorm:"column:Publisher;type:varchar(32)" json:"Publisher"`
	PubPlace    *string `gorm:"column:PubPlace;type:varchar(64)" json:"PubPlace"`
	PubDate     *string `gorm:"column:PubDate;type:varchar(32)" json:"PubDate"`
	Edition     *string `gorm:"column:Edition;type:varchar(32)" json:"Edition"`
	PrintNum    *string `gorm:"column:PrintNum;type:varchar(32)" json:"PrintNum"`
	Price       *string `gorm:"column:Price;type:varchar(32)" json:"Price"`
	Language    *string `gorm:"column:Language;type:varchar(32)" json:"Language"`
	Format      *string `gorm:"column:Format;type:varchar(32)" json:"Format"`
	Binding     *string `gorm:"column:Binding;type:varchar(32)" json:"Binding"`
	CatalogCode *string `gorm:"column:CatalogCode;type:varchar(32)" json:"CatalogCode"`
	Tags        *string `gorm:"column:Tags;type:varchar(64)" json:"Tags"`
	Summary     *string `gorm:"column:Summary;type:longtext" json:"Summary"`
	CIPData     *string `gorm:"column:CIPData;type:longtext" json:"CIPData"`
	Pages       *string `gorm:"column:Pages;type:varchar(32)" json:"Pages"`
	PrintCount  *string `gorm:"column:PrintCount;type:varchar(32)" json:"PrintCount"`
	ISBN10      *string `gorm:"column:ISBN10;type:varchar(32)" json:"ISBN10"`
	ISBN13      *string `gorm:"column:ISBN13;type:varchar(32)" json:"ISBN13"`
}

// TableName Rescipinfo's table name
func (*Rescipinfo) TableName() string {
	return TableNameRescipinfo
}
