package response

import "GINOWEN/models"

type LiblayerDto struct {
	models.Liblayer
	IsDeleted bool `gorm:"column:IsDeleted;type:bit(1);not null" json:"IsDeleted"`
	IsEnable  bool `gorm:"column:IsEnable;type:bit(1);not null" json:"IsEnable"`
}

type LibshelfDto struct {
	models.Libshelf
	IsDeleted  bool          `gorm:"column:IsDeleted;type:bit(1);not null" json:"IsDeleted"`
	IsEnable   bool          `gorm:"column:IsEnable;type:bit(1);not null" json:"IsEnable"`
	IsBosseyed bool          `gorm:"column:IsBosseyed;type:bit(1);not null;default:0" json:"IsBosseyed"`
	Layers     []LiblayerDto `json:"Layers"` // 层
}

type LibrowDto struct {
	models.Librow
	IsDeleted bool          `gorm:"column:IsDeleted;type:bit(1);not null" json:"IsDeleted"`
	Shelfs    []LibshelfDto `json:"Shelfs"` // 架
}

type QueryRowDto struct {
	Rows []LibrowDto `json:"Rows"` // 行
}
