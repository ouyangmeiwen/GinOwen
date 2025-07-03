package response

import (
	"GINOWEN/models"
	"GINOWEN/models/dto"
)

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
type DetailStatusListDto struct {
	LayerId      string `form:"LayerId" json:"LayerId"`
	LayerCode    string `form:"LayerCode" json:"LayerCode"`
	LayerName    string `form:"LayerName" json:"LayerName"`
	TaskStatus   int    `form:"TaskStatus" json:"TaskStatus"`
	ExceptionMsg string `form:"ExceptionMsg" json:"ExceptionMsg"`
	Remark       string `form:"Remark" json:"Remark"`
}
type PageDetailStatusListDto struct {
	TotalCount int                   `form:"TotalCount" json:"TotalCount"`
	Items      []DetailStatusListDto `form:"Items" json:"Items"`
	Detail     TaskDetail            `form:"TaskDetail" json:"TaskDetail"`
}
type InventoryMonthListDto struct {
	DateTime       string `form:"DateTime" json:"DateTime"`
	InventoryState int    `form:"InventoryState" json:"InventoryState"`
	Count          int    `form:"Count" json:"Count"`
}
type BooksNewIndexDto struct {
	TotalCount       int `form:"TotalCount" json:"TotalCount"`
	Seg_num          int `form:"seg_num" json:"seg_num"`
	Ocr_num          int `form:"ocr_num" json:"ocr_num"`
	Match_num        int `form:"match_num" json:"match_num"`
	Confidence_num   int `form:"confidence_num" json:"confidence_num"`
	Unconfidence_num int `form:"unconfidence_num" json:"unconfidence_num"`
}
type GetNotHitRankDto struct {
	dto.GetNotHitListDto
}
type GetFaultRankDto struct {
	LayerId   string `form:"LayerId" json:"LayerId"`
	LayerName string `form:"LayerName" json:"LayerName"`
	Count     int    `form:"Count" json:"Count"`
	Time      string `form:"Time" json:"Time"`
}
type BooksIndexDto struct {
	ItemsTotal              int    `form:"ItemsTotal" json:"ItemsTotal"`
	ItemLastYear            int    `form:"ItemLastYear" json:"ItemLastYear"`
	ItemCurrentYear         int    `form:"ItemCurrentYear" json:"ItemCurrentYear"`
	OffBooksTotal           int    `form:"OffBooksTotal" json:"OffBooksTotal"`
	OffBooksLastMonth       int    `form:"OffBooksLastMonth" json:"OffBooksLastMonth"`
	OffBookCurrentMonth     int    `form:"OffBookCurrentMonth" json:"OffBookCurrentMonth"`
	OnbooksTotal            int    `form:"OnbooksTotal" json:"OnbooksTotal"`
	OnBookCurrentMonth      int    `form:"OnBookCurrentMonth" json:"OnBookCurrentMonth"`
	OnbooksLastMonth        int    `form:"OnbooksLastMonth" json:"OnbooksLastMonth"`
	BookExceptionTotal      int    `form:"BookExceptionTotal" json:"BookExceptionTotal"`
	BookExceptionUpdateTime string `form:"BookExceptionUpdateTime" json:"BookExceptionUpdateTime"`
}
type BookRankIndexDto struct {
	Title     string `form:"Title" json:"Title"`
	Count     int    `form:"Count" json:"Count"`
	LastCount int    `form:"LastCount" json:"LastCount"`
	Index     int    `form:"Index" json:"Index"`
	CallNo    string `form:"CallNo" json:"CallNo"`
}
type InventoryFlyReadHisDto struct {
	CreateTime     string  `json:"createTime"`  //"2024-05-03 17:44:58",//创建时间
	UpdateTime     string  `json:"updateTime"`  //"2024-05-03 17:47:58",//更新时间
	WorkId         string  `json:"workId"`      // "706441161815887872",//盘点id
	CheckType      string  `json:"checkType"`   //"REALTIME",//任务类型
	CheckAll       bool    `json:"checkAll"`    //是否全馆
	DeviceType     string  `json:"deviceType"`  //"CAMERA",//盘点设备类型
	Status         string  `json:"status"`      //"EXPORTED",//盘点任务状态 INIT-创建 COLLECTING-采集 COMPUTED-盘点 EXPORTED-盘点完成 FAILED-失败
	TaskAllNum     int     `json:"taskAllNum"`  //盘点总个数
	FinishNum      int     `json:"finishNum"`   //完成个数
	BookCount      int     `json:"bookCount"`   //已盘点图书
	AssertCount    int     `json:"assertCount"` //馆藏总数
	TotalNums      int     `json:"TotalNums"`
	OnAndFaultNums int     `json:"OnAndFaultNums"`
	TotalTime      float64 `json:"TotalTime"`
	OnNums         int     `json:"OnNums"`
	FaultNums      int     `json:"FaultNums"`
	OffNums        int     `json:"OffNums"`
}
type InventoryFlyReadIndexDto struct {
	Current InventoryFlyReadHisDto `form:"Current" json:"Current"`
	Last    InventoryFlyReadHisDto `form:"Last" json:"Last"`
}
type StructDto struct {
	models.Libstruct
	Children []StructDto `form:"Children" json:"Children"`
}
