package response

import "GINOWEN/models/dto"

type HelloResp struct {
	Message string `json:"message"` // 响应消息
	Status  int    `json:"status"`  // 响应状态码
	Success bool   `json:"success"` // 是否成功
}
type GetFlyTokenInputResp struct {
	AccessToken string `json:"accessToken"` // 飞阅访问令牌
}
type UploadLibItemResp struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type UploadTenantDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type UploadStructDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type UploadLibItemLocDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type UploadRowDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type InventoryDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type GetRobotRouterlistDto struct {
	Items []dto.FlyRouter `form:"items" json:"items"` // 是否成功
}
type InventoryHisDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
	dto.InventoryHisRespObj
}
type InventoryListDto struct {
	List []dto.InventoryHisRespObj `json:"list"` // 盘点列表
}
