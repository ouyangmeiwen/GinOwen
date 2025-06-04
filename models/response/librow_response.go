package response

import "GINOWEN/models"

type LiblayerDto struct {
	models.Liblayer
}

type LibshelfDto struct {
	models.Libshelf
	Layers []LiblayerDto `json:"Layers"` // 层
}

type LibrowDto struct {
	models.Librow
	Shelfs []LibshelfDto `json:"Shelfs"` // 架
}

type QueryRowDto struct {
	Rows []LibrowDto `json:"Rows"` // 行
}
