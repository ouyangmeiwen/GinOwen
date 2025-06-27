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
type SetBussinessDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type GetInventorySetDto struct {
	SysStartTime       string `json:"sysStartTime"`       // 系统开始时间
	SysEndTime         string `json:"sysEndTime"`         // 系统结束时间
	IsEnable           bool   `json:"isEnable"`           // 是否启用
	InventoryStartDate string `json:"inventoryStartDate"` // 盘点开始时间
	InventoryEndDate   string `json:"inventoryEndDate"`   // 盘点结束时间
	Interval           int    `json:"interval"`           // 盘点间隔时间
	DeviceType         string `json:"deviceType"`         // 设备类型
}
type GetEnableRowDto struct {
	dto.GetEnableRowDto
}
type GetRobotlistDto struct {
	dto.RobotlistDto
}
type GetCaseImgsDto struct {
	dto.GetCaseImgsDto
}
type GetOcrImgsDto struct {
	dto.GetOcrImgsDto
}
type InventorySetDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type CreatWorkDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type UpdateWorkDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type WorkListDto struct {
	WorkName          string `form:"WorkName" json:"WorkName"`
	WorkId            string `form:"WorkId" json:"WorkId"`
	WorkTime          string `form:"WorkTime" json:"WorkTime"`
	WorkStarTime      string `form:"WorkStarTime" json:"WorkStarTime"`
	WorkEndTime       string `form:"WorkEndTime" json:"WorkEndTime"`
	CreateTime        string `form:"CreateTime" json:"CreateTime"`
	TaskStatus        int    `form:"TaskStatus" json:"TaskStatus"`
	TriggerSatus      int    `form:"TriggerSatus" json:"TriggerSatus"`
	Comment           string `form:"Comment" json:"Comment"`
	DetailsCount      int    `form:"DetailsCount" json:"DetailsCount"`
	ExceptionMsgCount int    `form:"ExceptionMsgCount" json:"ExceptionMsgCount"`
	ExceptionMsg      string `form:"ExceptionMsg" json:"ExceptionMsg"`
}
type PageWorkListDto struct {
	TotalCount     int           `form:"TotalCount" json:"TotalCount"`
	Items          []WorkListDto `form:"Items" json:"Items"`
	InitWorkCount  int           `form:"InitWorkCount" json:"InitWorkCount"`
	InventoryCount int           `form:"InventoryCount" json:"InventoryCount"`
	SucessCount    int           `form:"SucessCount" json:"SucessCount"`
	FailtureCount  int           `form:"FailtureCount" json:"FailtureCount"`
}
type DeleteWorkDto struct {
	Success bool `form:"success" json:"success"` // 是否成功
}
type DetailListDto struct {
	LayerId        string `form:"LayerId" json:"LayerId"`
	LayerCode      string `form:"LayerCode" json:"LayerCode"`
	LayerName      string `form:"LayerName" json:"LayerName"`
	Title          string `form:"Title" json:"Title"`
	ISBN           string `form:"ISBN" json:"ISBN"`
	Author         string `form:"Author" json:"Author"`
	Publisher      string `form:"Publisher" json:"Publisher"`
	CallNo         string `form:"CallNo" json:"CallNo"`
	Barcode        string `form:"Barcode" json:"Barcode"`
	LocationName   string `form:"LocationName" json:"LocationName"`
	InventoryState int    `form:"InventoryState" json:"InventoryState"`
	LocLayerName   string `form:"LocLayerName" json:"LocLayerName"`
	LocLayerId     string `form:"LocLayerId" json:"LocLayerId"`
	LocLayerCode   string `form:"LocLayerCode" json:"LocLayerCode"`
	Remark         string `form:"Remark" json:"Remark"`
}

type TaskDetail struct {
	DeviceType      string `form:"DeviceType" json:"DeviceType"`
	RobotId         string `form:"RobotId" json:"RobotId"`
	RobotName       string `form:"RobotName" json:"RobotName"`
	RobotRouterId   string `form:"RobotRouterId" json:"RobotRouterId"`
	RobotRouterName string `form:"RobotRouterName" json:"RobotRouterName"`
	WorkName        string `form:"WorkName" json:"WorkName"`
	WorkId          string `form:"WorkId" json:"WorkId"`
	TriggerSatus    int    `form:"TriggerSatus" json:"TriggerSatus"`
	WorkTime        string `form:"WorkTime" json:"WorkTime"`
	WorkStarTime    string `form:"WorkStarTime" json:"WorkStarTime"`
	WorkEndTime     string `form:"WorkEndTime" json:"WorkEndTime"`
	WorkLayerCount  int    `form:"WorkLayerCount" json:"WorkLayerCount"`
}
type PageDetailListDto struct {
	TotalCount int             `form:"TotalCount" json:"TotalCount"`
	Items      []DetailListDto `form:"Items" json:"Items"`
	Detail     TaskDetail      `form:"TaskDetail" json:"TaskDetail"`
}
