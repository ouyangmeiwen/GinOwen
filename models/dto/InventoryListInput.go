package dto

type InventoryListObj struct {
	Offset     int    `json:"offset"`     //分页偏移量 0
	PageSize   int    `json:"pageSize"`   //分页大小 20
	DeviceType string `json:"deviceType"` //CAMERA
	StartTime  string `json:"startTime"`  //"2024-05-01",
	EndTime    string `json:"endTime"`    //"2024-05-01",
}

type InventoryListInput struct {
	Obj   InventoryListObj `json:"obj"`
	Token string           `json:"token"`
}
type InventoryListDtoObj struct {
	List []InventoryHisRespObj `json:"list"`
}
type InventoryListDto struct {
	FlyReadBaseResp
	Obj InventoryListDtoObj `json:"obj"` //盘点列表
}
