package models

import (
	"database/sql"
	"time"
)

type LibItem struct {
	Id                   string         `json:"Id" gorm:"primaryKey;column:Id;comment:Id"`
	CreationTime         time.Time      `json:"CreationTime" gorm:"column:CreationTime;comment:CreationTime:"`
	CreatorUserId        sql.NullInt64  `json:"CreatorUserId" gorm:"column:CreatorUserId;comment:CreatorUserId:"`
	LastModificationTime sql.NullTime   `json:"LastModificationTime" gorm:"column:LastModificationTime;comment:LastModificationTime:"`
	LastModifierUserId   sql.NullInt64  `json:"LastModifierUserId" gorm:"column:LastModifierUserId;comment:LastModifierUserId:"`
	IsDeleted            []uint8        `json:"IsDeleted" gorm:"column:IsDeleted;comment:IsDeleted:"` //注意不可用interface{}
	DeleterUserId        sql.NullInt64  `json:"DeleterUserId" gorm:"column:DeleterUserId;comment:DeleterUserId:"`
	DeletionTime         sql.NullTime   `json:"DeletionTime" gorm:"column:DeletionTime;comment:DeletionTime:"`
	InfoId               sql.NullString `json:"InfoId" gorm:"column:InfoId;comment:InfoId"`
	Title                string         `json:"Title" gorm:"column:Title;comment:Title"`
	Author               sql.NullString `json:"Author" gorm:"column:Author;comment:Author"`
	Barcode              string         `json:"Barcode" gorm:"column:Barcode;comment:Barcode"`
	IsEnable             []uint8        `json:"IsEnable" gorm:"column:IsEnable;comment:IsEnable"` //注意不可用interface{}
	CallNo               sql.NullString `json:"CallNo" gorm:"column:CallNo;comment:CallNo"`
	PreCallNo            sql.NullString `json:"PreCallNo" gorm:"column:PreCallNo;comment:PreCallNo"`
	CatalogCode          sql.NullString `json:"CatalogCode" gorm:"column:CatalogCode;comment:CatalogCode"`
	ItemState            uint8          `json:"ItemState" gorm:"column:ItemState;comment:ItemState"`
	PressmarkId          sql.NullString `json:"PressmarkId" gorm:"column:PressmarkId;comment:PressmarkId"`
	PressmarkName        sql.NullString `json:"PressmarkName" gorm:"column:PressmarkName;comment:PressmarkName"`
	LocationId           sql.NullString `json:"LocationId" gorm:"column:LocationId;comment:LocationId"`
	LocationName         sql.NullString `json:"LocationName" gorm:"column:LocationName;comment:LocationName"`
	BookBarcode          sql.NullString `json:"BookBarcode" gorm:"column:BookBarcode;comment:BookBarcode"`
	ISBN                 sql.NullString `json:"ISBN" gorm:"ISBN;column:ISBN;comment:ISBN"`
	PubNo                sql.NullInt16  `json:"PubNo" gorm:"column:PubNo;comment:PubNo"`
	Publisher            sql.NullString `json:"Publisher" gorm:"column:Publisher;comment:Publisher"`
	PubDate              sql.NullString `json:"PubDate" gorm:"column:PubDate;comment:PubDate"`
	Price                sql.NullString `json:"Price" gorm:"column:Price;comment:Price"`
	Pages                sql.NullString `json:"Pages" gorm:"column:Pages;comment:Pages"`
	Summary              sql.NullString `json:"Summary" gorm:"column:Summary;comment:Summary"`
	ItemType             uint8          `json:"ItemType" gorm:"column:ItemType;comment:ItemType"`
	Remark               sql.NullString `json:"Remark" gorm:"column:Remark;comment:Remark"`
	OriginType           uint8          `json:"OriginType" gorm:"column:OriginType;comment:OriginType"`
	CreateType           uint8          `json:"CreateType" gorm:"column:CreateType;comment:CreateType"`
	TenantId             int            `json:"TenantId" gorm:"column:TenantId;comment:TenantId"`
}

func (LibItem) TableName() string {
	return "libitem"
}
