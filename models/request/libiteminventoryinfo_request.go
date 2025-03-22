package request

import (
	"GINOWEN/models"
	"GINOWEN/utils"
)

// get请求必须指定json form
type GetInventoryInput struct {
	utils.PagedResultRequest
	Barcode      string `json:"barcode" form:"barcode"`
	Title        string `json:"title" form:"title"`
	CallNo       string `json:"callNo" form:"callNo"`
	ISBN         string `json:"isbn" form:"isbn"`
	LayerCode    string `json:"layerCode" form:"layerCode"`
	LayerName    string `json:"layerName" form:"layerName"`
	LocLayerCode string `json:"locLayerCode" form:"locLayerCode"`
	LocLayerName string `json:"locLayerName" form:"locLayerName"`
}

type CreateInventoryInput struct {
	Item models.Libiteminventoryinfo
}

type UpdateInventoryInput struct {
	Item models.Libiteminventoryinfo
}

// delete请求必须指定json form
type DeleteInventoryInput struct {
	Id string `json:"id" form:"id"`
}
